package main

import (  
    "io/ioutil"
	"os"
	"strings"
	"time"
	"fmt"
)  

//Read from table dumped file
func Read(table string) []string {
	dat, err := ioutil.ReadFile("./data/"+table+".txt")
	check(err) 
	res := strings.Split(string(dat),"\n")	
	return removeLastSpace(res)
}

func removeLastSpace(data []string) []string{
	if len(data[len(data)-1]) == 0 {
		return data[:len(data)-1]
	} 
	return data
}

//Dumps dump all tables with update time
func Dumps(){
	tables := LoadTablesWithUpdate()
	jobs := make(chan string, 1000)
	results := make(chan int, 1000)
	for w := 1; w <= 8; w++ {
        go dumpWorker(w,jobs,results)
	}
	for _,table := range tables {
		if len(table) > 1 {
			jobs <- table			
		}
	}
	close(jobs)
	waitDumpResult(tables,results) 
}
 
func dumpWorker(id int, jobs <-chan string, result chan<- int) {
    for table := range jobs {
		start := time.Now().UnixNano() / 1000000 
		res := Fetch(table)	
		Dump(table,res)
		end := time.Now().UnixNano() / 1000000 
		fmt.Println(table ," spend time: " ,(end - start), "ms @worker:",id)
		result <- 1			
    }
}
func waitDumpResult(tables []string, results chan int) {			
	for i,table := range tables {
		if len(table) > 1 {
			<-results	
			fmt.Println(i,"/",len(tables))
		} 
	}	 
}

//Dump data to table named file
func Dump(table string, res []string) {
	os.Mkdir("data",0777)
	f, err := os.Create("./data/"+table+".txt")
	f.Seek(0,0)
	check(err)
	defer f.Close()
	for _,l := range res {			
		f.WriteString(l + "\n")
	}	
	f.Sync() 
}

