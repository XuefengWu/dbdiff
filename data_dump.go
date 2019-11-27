package main

import (  
    //"io/ioutil"
	//"os"
	"database/sql"
	"strings"
	"time"
	"fmt"
)  

//Read from table dumped file
func Read(baselines map[string]string,table string) []string {
	//dat, err := ioutil.ReadFile("./data/"+table+".txt")
	//check(err) 
	l := baselines[table]
	res := strings.Split(l,"\n")	
	return removeLastSpace(res)
}

func removeLastSpace(data []string) []string{
	if len(data[len(data)-1]) == 0 {
		return data[:len(data)-1]
	} 
	return data
}

//Dumps dump all tables with update time
func Dumps(remainDay int,db *sql.DB,tables []string, results chan<- map[string]string) {
	
	jobs := make(chan string, 1000)
	//results := make(chan map[string]string, 1000)

	//read from db	


	for w := 1; w <= 32; w++ {
        go dumpWorker(w,remainDay,db,jobs,results)
	}
	for _,table := range tables {
		if len(table) > 1 {
			jobs <- table			
		}
	}
	close(jobs)
	//ret := waitDumpResult(tables,results) 
	//close(results)
	//return ret
}
 
func dumpWorker(id int,remainDay int,db *sql.DB, jobs <-chan string, result chan<- map[string]string) {
    for table := range jobs {
		start := time.Now().UnixNano() / 1000000 

		res := Fetch(table,remainDay,db)			
		r := dump(table,res)
		m := map[string]string{table:r}
		result <- m			
		end := time.Now().UnixNano() / 1000000 
		spend := (end - start)/1000
		if spend > 3 {
			fmt.Println(table ," spend time: " ,spend, "s @worker:",id)
		}		
    }
}

//WaitDumpResult Wait Result channel finished
func WaitDumpResult(tables []string, results chan map[string]string) map[string]string{	
	records := make(map[string]string)		
	for _,table := range tables {
		if len(table) > 1 {
			r := <- results	
			for t, l := range r {
				records[t] = l
			}			
			//fmt.Println(i+1,"/",len(tables))
		} 
	}	
	fmt.Println("WaitDumpResult finished:",len(tables))
	return records 
}

//Dump data to table named file
func dump(table string, res []string) string{
	// os.Mkdir("data",0777)
	// f, err := os.Create("./data/"+table+".txt")
	// f.Seek(0,0)
	// check(err)
	// defer f.Close()
	// for _,l := range res {			
	// 	f.WriteString(l + "\n")
	// }	
	// f.Sync() 
	//if len(res) > 0 {
	//	fmt.Println(table,":",strings.Join(res,"\n"))
	//}
	return strings.Join(res,"\n")
}

