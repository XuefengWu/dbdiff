package main

import ( 	
	"fmt"  
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//DumpTablesWithUpdateTime dump tables with update time
func DumpTablesWithUpdateTime(connString string) {
	dumpTables(connString)
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

	//tablesWithUpdateTime := make([]string,0)
	tablesWithUpdateTime := make(map[string]struct{})
	for t,tcs := range tables { 
		for _,c := range tcs { 
			if "UPDATE_TIME" == c {
				tablesWithUpdateTime[t] = struct{}{}				
			}			
		}
	} 
	fmt.Println(tablesWithUpdateTime)
	f, err := os.Create("./data/tables_update.txt")
	check(err)
	defer f.Close()
	for t,_ := range tablesWithUpdateTime {
		f.WriteString(t + "\n")
	}
}

func dumpTables(connString string) {
	tables := LoadTables(connString)
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
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err) 
	fmt.Println("LoadTablesWithUpdate current directory: ",dir)
	dat, err := ioutil.ReadFile("./data/tables_update.txt")
	check(err) 
	lines := strings.Split(string(dat),"\n")
	tables := make([]string,0)
	for _,l := range lines {
		if !strings.HasPrefix(l, "-") {
			tables = append(tables,l)
		}		
	}
	return tables
}