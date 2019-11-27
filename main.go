package main

import ( 	
	"fmt" 
	"os"
	"time"	 
	"log"
	"net/http"
)
/**
* dump_tables "dev/dev@127.0.0.1/dev"  
* dump "dev/dev@127.0.0.1/dev"  1
* diff "dev/dev@127.0.0.1/dev"  1
* http 
*   -- POST diff "dev/dev@127.0.0.1/dev" "dev/dev@127.0.0.2/dev"  1
**/
func main() {	
	args := os.Args
	act := args[1]	
	start := time.Now().UnixNano() / 1000000 
	if act == "dump" {
		//remainDay := RemainDayConfig()
		//connString := ConnStringConfig()
		//Dumps(remainDay,connString)
	}else if act == "dump_tables"{
		//connString := ConnStringConfig()
		//DumpTablesWithUpdateTime(connString)		
	} else if act == "diff" {
		//remainDay := RemainDayConfig()
		//connString := ConnStringConfig()
		//ls := make(map[string]string)
		//res := Diffs(remainDay,connString,ls)
		//DumpHTMLReport(res)
		//tables := LoadTablesWithUpdate()
		//DumpJunitReport(tables, res)
	} else {
		http.HandleFunc("/diff", handleDiff)		 
		http.HandleFunc("/report", reportHandler)
		http.HandleFunc("/", handler)
		
		fmt.Println("running on : " , 8080)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
	end := time.Now().UnixNano() / 1000000 
	fmt.Println("spend time: " , (end - start) / 1000," s")
}

