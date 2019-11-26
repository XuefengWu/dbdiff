package main

import (    
	"database/sql"	
	_ "gopkg.in/goracle.v2"
) 

//CreateConn create connection for oracle DB
func CreateConn(connString string) (*sql.DB, error) { 	
	db, err := sql.Open("goracle", connString)	
	check(err)
	err = db.Ping()
	check(err)
	return db, err
}

