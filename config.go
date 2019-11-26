package main

import ( 
	"os" 
	"strconv" 
)

func RemainDayConfig() int {
	args := os.Args	
	if len(args) > 3 {
		remain := args[3]
		i64, err := strconv.ParseInt(remain, 10, 32)
		check(err)
		return int(i64)
	} 
	return 7
}

func ConnStringConfig() string {
	args := os.Args	
	if len(args) > 2 {
		conn := args[2]
		return conn
	} 
	return "dev/dev@127.0.0.1/dev_db"
}