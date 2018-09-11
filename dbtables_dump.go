package main

import ( 	
	"fmt"  
	"os"
	"io/ioutil"
	"strings"
)

//DumpTablesWithUpdateTime dump tables with update time
func DumpTablesWithUpdateTime() {
	dat, err := ioutil.ReadFile("./data/tables.txt")
	check(err) 
	lines := strings.Split(string(dat),"\n")
	tables := make(map[string][]string)
	var t string
	for _,l := range lines {
		if !strings.HasPrefix(l,"\t") {
			t = l
			tables[l] = make([]string,0)
		} else {
			tables[t] = append(tables[t],strings.TrimSpace(l))
		}
	}

	tablesWithUpdateTime := make([]string,0)
	for t,tcs := range tables { 
		for _,c := range tcs { 
			if "UPDATE_TIME" == c {
				tablesWithUpdateTime = append(tablesWithUpdateTime,t)
			}			
		}
	} 
	fmt.Println(tablesWithUpdateTime)
	f, err := os.Create("./data/tables_update.txt")
	check(err)
	defer f.Close()
	for _,t := range tablesWithUpdateTime {
		f.WriteString(t + "\n")
	}
}

func dumpTables() {
	tables := LoadTables()
	f, err := os.Create("./data/tables.txt")
	check(err)
	defer f.Close()
	for t,tcs := range tables {
		f.WriteString(t + "\n")
		for _,c := range tcs {
			f.WriteString("\t" + c + "\n")
		}
	} 
	f.Sync()
	fmt.Println("Finished.")
}
 
//LoadTablesWithUpdate tables with update time
func LoadTablesWithUpdate() []string {
	dat, err := ioutil.ReadFile("./data/tables_update.txt")
	check(err) 
	lines := strings.Split(string(dat),"\n")
	tables := make([]string,0)
	for _,l := range lines {
		tables = append(tables,l)
	}
	return tables
}