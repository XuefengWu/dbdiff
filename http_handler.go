package main

import ( 	
	"fmt"  
	"time"	
	"strconv" 
	"net/http"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handleDiff(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form 
	baseline := form["baseline"][0]
	target := form["target"][0]
	remain := form["remain"][0]
	i64, err := strconv.ParseInt(remain, 10, 32)
	check(err)
	remainDay := int(i64)

	start := time.Now().UnixNano() / 1000000
	baselines := Dumps(remainDay,baseline)
	res := Diffs(remainDay,target,baselines)
	DumpHTMLReport(res)
	jsonString, err := json.Marshal(res)
	check(err)
	end := time.Now().UnixNano() / 1000000 
	fmt.Println("Diff spend time: " , (end - start) / 1000," s")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if(len(res) > 0) {
		w.WriteHeader(417) //Expectation Failed
	} else {
		w.WriteHeader(200)
	}
	w.Write(jsonString)	
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    http.ServeFile(w, r, "./data/report.html")
}