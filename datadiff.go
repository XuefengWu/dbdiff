package main

import (  	
	"fmt"  
	"github.com/sergi/go-diff/diffmatchpatch"			
)

//Diffs all tables with update time
func Diffs(baselines map[string]string,targets map[string]string,tables []string)  map[string][]string {	
	diffs := make(map[string][]string)
	for _,table := range tables {
		if len(table) > 1 {
			ds := diff(baselines,targets,table)
			for t,v := range ds {
			if len(v) > 0 {					
				diffs[t] = v
			} 
			}
		}
	}
	return diffs
}

func waitDiffResult(tables []string, results chan map[string][]string) map[string][]string {		
	diffs := make(map[string][]string)
	for _,table := range tables {
		if len(table) > 1 {
			ds := <-results	
			for t,v := range ds {
			if len(v) > 0 {					
				diffs[t] = v
			}
			//fmt.Println(i+1,"/",len(tables))
			}
		}
	}	
	fmt.Println("datadiff waitDiffResult finished")
	return diffs
}



func worker(baselines map[string]string,targets map[string]string, results chan map[string][]string,jobs chan string) {
    for table := range jobs { 
		diffs := diff(baselines,targets,table)
		results <- diffs		
    }
}

func diff(baselines map[string]string,targets map[string]string,table string) map[string][]string {
	res := Read(targets,table)  
	txts := Read(baselines,table)  
	diffs := diffText(table,txts,res)	
	return map[string][]string{table:diffs}
}

//Diff : tell different between
func diffText(table string,texts1 []string, texts2 []string) []string { 
	//fmt.Println("Diff: ", table)
	//fmt.Println("target record size: ",len(texts1))
	//fmt.Println("acutal record size: ",len(texts2))
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
 
