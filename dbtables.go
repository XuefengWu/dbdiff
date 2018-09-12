package main

import (
	"fmt"  
	"time" 	
	"context"	
	"database/sql"	
	_ "gopkg.in/goracle.v2"
)

//LoadTables from DB for table: query table name
func LoadTables() map[string][]string {
	//read from db
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()
 	db,err := CreateConn()	 
	 if err != nil { 
		return nil
	}
	defer db.Close()
	tables := queryTables(ctx,db) 	
	res := make(map[string][]string)
	for _,t := range tables {		
		columns := queryColumns(ctx,db,t)
		fmt.Println(t," ",len(columns))
		res[t] = columns
	}
	return res	  
}

func queryTables(ctx context.Context,db *sql.DB) []string {
	q := "SELECT table_name FROM dba_tables WHERE owner=:db_owner"
	//MARK 
	rows, err :=  db.QueryContext(ctx,q,sql.Named("db_owner", "RELEASETEST"))	 
	return scanOneColumns(rows,err)
}

func queryColumns(ctx context.Context,db *sql.DB,table string) []string {
	q := "SELECT column_name FROM all_tab_cols WHERE table_name = :table_name"
	rows, err :=  db.QueryContext(ctx,q,sql.Named("table_name", table))	 
	return scanOneColumns(rows,err)	 
}
  
func scanOneColumns(rows *sql.Rows,err error) []string {			 
	if err != nil {
		fmt.Println("QueryContext: ",err)
		return nil
	}
	res := make([]string,0)
	for rows.Next() { 
		var r string				
		if err := rows.Scan(&r); err != nil {
			fmt.Println("Scan: ",err)
		}	 				
		res = append(res,r)		
	}
	return res
}