package main

import ( 	
	"fmt" 
	"os"
	"time"	
)

func main() {	
	args := os.Args
	act := args[1]	
	start := time.Now().UnixNano() / 1000000 
	if act == "dump" {
		Dumps()
	}else if act == "dump_tables"{
		DumpTablesWithUpdateTime()		
	} else {
		res := Diffs()
		DumpHTMLReport(res)
		tables := LoadTablesWithUpdate()
		DumpJunitReport(tables, res)
	}
	end := time.Now().UnixNano() / 1000000 
	fmt.Println("spend time: " , (end - start) / 1000," s")
}