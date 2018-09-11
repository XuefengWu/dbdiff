package main

import (
	"fmt"  
	"time"  
	"context"
	"os"
	"regexp"
	"strings"
	"strconv"
	"database/sql"	
	_ "gopkg.in/goracle.v2"
)

//Fetch from DB for table: query table name
func Fetch(table string) []string {
	if len(table) < 1 {
		return nil
	}
	//read from db
	ctx, cancel := context.WithTimeout(context.Background(), 1800*time.Second)
	defer cancel()
 	db,err := CreateConn()	 
	 if err != nil { 
		return nil
	}
	defer db.Close()
	return queryRecords(ctx,db,table) 	  
} 
 
func queryTablePrimaryKey(ctx context.Context,db *sql.DB, table string) []string {
	q := fmt.Sprintf(`SELECT cols.column_name
	FROM all_constraints cons, all_cons_columns cols
	WHERE cols.table_name = '%s'
	AND cons.constraint_type = 'P'
	AND cons.owner = 'RELEASETEST'
	AND cons.constraint_name = cols.constraint_name
	AND cons.owner = cols.owner`,table)
	
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		fmt.Println(err)
		return []string{"1"}
	}
	res := make([]string,0)
	for rows.Next() { 			
		var v string	
		if err := rows.Scan(&v); err != nil {
			fmt.Println(err)
		}	 					
		res = append(res,v)
	}
	return res
}
func queryRecords(ctx context.Context,db *sql.DB, table string) []string {
	const remainDay = 7
	pks := queryTablePrimaryKey(ctx,db,table)
	orders := strings.Join(pks,",")
	q := fmt.Sprintf("SELECT * FROM %s where update_time > trunc((SYSDATE - :remain_day)) ORDER BY %s", table, orders)
	
	rows, err := db.QueryContext(ctx, q, sql.Named("remain_day", remainDay))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cols,_ := rows.ColumnTypes()
	  
	vs := make([]interface{},len(cols))
	dest := make([]interface{},0)	
	for i := range vs { 
		dest = append(dest,&vs[i])
	}	
	res := make([]string,0)
	for rows.Next() { 				
		if err := rows.Scan(dest...); err != nil {
			fmt.Println(err)
		}		
		var r string		
		for _, v := range vs { 			
			if !isClob(v) {
				r += fmt.Sprint(trimTime(v)," ")
			}			
		}					
		res = append(res,trimRecord(r))
	}
	return res
}

func isClob(line interface{}) bool {
	record := fmt.Sprint(line)
	res,_ := regexp.MatchString("&{.*}", record)
	return res
}

func trimTime(v interface{}) string {
	str := fmt.Sprint(v)
	timeMatched,_ := regexp.MatchString("\\d\\d\\d\\d-\\d\\d-\\d\\d .*[C|T]", str)
	if timeMatched {
		str = str[:7]
	}
	return str
}
func trimRecord(line string) string {
	str := strings.Replace(line, "\n", "_", -1)
	str = strings.Replace(str, "\r", "_", -1)
	str = strings.Replace(str, "\t", "_", -1)
	str = strings.Replace(str, "   ", "_", -1)
	return str
}

func remainDayConfig() int {
	args := os.Args	
	if len(args) > 3 {
		remain := args[3]
		i64, err := strconv.ParseInt(remain, 10, 32)
		check(err)
		return int(i64)
	} 
	return 7
}