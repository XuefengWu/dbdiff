package main

import (  	
	"fmt" 
	"time"
	"github.com/sergi/go-diff/diffmatchpatch"			
)

//Diffs all tables with update time
func Diffs()  map[string][]string {
	tables := LoadTablesWithUpdate()
	jobs := make(chan string, 1000)
	results := make(chan map[string][]string, 1000)
	for w := 1; w <= 4; w++ {
        go worker(w,jobs,results)
    }
	for _,table := range tables {
		if len(table) > 1 {
			jobs <- table			
		}
	}
	close(jobs)
	diffResult := waitDiffResult(tables,results)
	return diffResult
}

func waitDiffResult(tables []string, results chan map[string][]string) map[string][]string {		
	diffs := make(map[string][]string)
	for i,table := range tables {
		if len(table) > 1 {
			ds := <-results	
			for t,v := range ds {
			if len(v) > 0 {					
				diffs[t] = v
		}
			fmt.Println(i,"/",len(tables))
		}
	}
	}	
	return diffs
}


func worker(id int, jobs <-chan string, results chan<- map[string][]string) {
    for table := range jobs {
		start := time.Now().UnixNano() / 1000000 
		diffs := diff(table)
		end := time.Now().UnixNano() / 1000000 
		fmt.Println(table ,"diffs:",len(diffs)," spend time: " ,(end - start), "ms @worker:",id)
		results <- diffs		
    }
}

func diff(table string) map[string][]string {
	res := Fetch(table)	 
	txts := Read(table)  
	diffs := diffText(table,txts,res)	
	return map[string][]string{table:diffs}
}

//Diff : tell different between
func diffText(table string,texts1 []string, texts2 []string) []string { 
	fmt.Println("Diff: ", table)
	fmt.Println("target record size: ",len(texts1))
	fmt.Println("acutal record size: ",len(texts2))
	if len(texts1) != len(texts2) {
		res := make([]string,0)
		diffMessage := fmt.Sprintf("size of records is different %d != %d.",len(texts1),len(texts2))
		diffMessage += fmt.Sprintf(" target record size: %d, ",len(texts1))
		diffMessage += fmt.Sprintf("acutal record size: %d.",len(texts2))
		res = append(res,diffMessage)
		return res
	}
	dmp := diffmatchpatch.New()
	diffs := make([]string,0)
	for i,v := range texts1 { 
		diff := dmp.DiffMain(texts2[i],v, false)			
		if len(diff) > 1 {			
			diffs = append(diffs,dmp.DiffPrettyHtml(diff))			
		} 
	}
	return diffs 
}
 
