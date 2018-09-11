package main

import (   
	"os" 			
)
 
//DumpHTMLReport dump report as Html
func DumpHTMLReport(diffs map[string][]string)  {
	f, err := os.Create("./data/report.html") 
	check(err)
	defer f.Close()	  
	f.WriteString(diffs2HTML(diffs))
	f.Sync()  
}

func diffs2HTML(diffs map[string][]string) string {
	html := "<html>\n"
	html += "<head>\n"
	html += "\t<title>Diff Report</title>\n"
	html += "</head>\n"
	html += "<body>\n"
	for t,v := range diffs {
		if len(v) > 0 {					
			html += "\t<h3>"+ t + "</h3>\n"
			html += "\t<div>\n\t\t<div>\n"			
			for _,l := range v {						
				html += "\t\t\t<div>"+l + "</div>\n"				
			}
			html += "\t\t</div>\n\t</div>\n"
	}			 
	}
	html += "</body>\n"
	html += "</html>"
	return html
}