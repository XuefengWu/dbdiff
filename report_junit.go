package main

import (   
	"encoding/xml"	
	"os" 			
	"sort"	
)

//Failure junit Failure
type Failure struct {
	XMLName 	xml.Name	`xml:"failure"`
	Message 	string  	`xml:"message"`
	Type 		string		`xml:"type,attr"`
}

//TestCase junit TestCase
type TestCase struct {
	XMLName 	xml.Name	`xml:"testcase"`
	ClassName 	string 		`xml:"classname,attr"`
	Name 	string 			`xml:"name,attr"`
	Failures []Failure
}

//TestSuite junit TestSuite
type TestSuite struct {
	XMLName	 	xml.Name 		`xml:"testsuite"`
	TestsSize 	int				`xml:"tests,attr"`
	TestCases 	[]TestCase
}

//TestSuites junit TestSuites
type TestSuites struct {
	XMLName	 	xml.Name 		`xml:"testsuites"`
	Failures 	int				`xml:"failures,attr"`
	Tests 		int				`xml:"tests,attr"`
	TestSuites	[]TestSuite
}

//DumpJunitReport dump junit report
func DumpJunitReport(tables []string, diffs map[string][]string) {	
	output, err := diffs2JunitXML(removeSpace(tables), diffs)
	check(err)
	f, err := os.Create("./data/report.xml")
	check(err)
	f.Write(output)	
}

func removeSpace(tables []string) []string{
	res := make([]string,0)
	for _,t := range tables {
		if len(t) > 1 {
			res = append(res,t)
		}
	}
	return res
}

func diffs2JunitXML(tables []string, diffs map[string][]string) ([]byte, error) {
	sort.Strings(tables)
	failuresSize := len(diffs)
	testSize := len(tables)

	testCases := make([]TestCase,0)
	for _,table := range tables {
		value, ok := diffs[table]
		var tc TestCase
        if ok {
			failure := failureNode(value)
			tc = TestCase{ClassName:table, Name:table,Failures:[]Failure{failure}}
        } else {
            tc = TestCase{ClassName:table, Name:table}
		}
		testCases = append(testCases,tc)
	}
	ts  := TestSuite{TestsSize:testSize,TestCases:testCases}
	tss := TestSuites{Failures: failuresSize, Tests:testSize,TestSuites:[]TestSuite{ts}} 

	return xml.MarshalIndent(tss, "  ", "    ")	
}

func failureNode(diffs []string) Failure {
	message := "<![CDATA["
	for _,dif := range diffs {
		message += dif + "\n"
	}
	message += "]]>"
	return Failure{Message: message, Type:"Different"}
}
