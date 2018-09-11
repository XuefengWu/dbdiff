package main

import (   
	"os"
	"database/sql"	
	_ "gopkg.in/goracle.v2"
) 

//CreateConn create connection for oracle DB
func CreateConn() (*sql.DB, error) { 
	connString := connStringConfig()
	db, err := sql.Open("goracle", connString)	
	check(err)
	err = db.Ping()
	check(err)
	return db, err
}

func connStringConfig() string {
	args := os.Args	
	if len(args) > 2 {
		conn := args[2]
		return conn
	} 
	return "dev/dev@127.0.0.1/dev_db"
}