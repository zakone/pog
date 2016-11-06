//go run issueServer.go &
//localhost:8000/?repo=golang/go
package main

import (
	"./github"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const templ = `<h1>{{.TotalCount}} issues:</h1>
<table>
<tr style='text-align: left'>
 <th>#</th>
 <th>State</th>
 <th>User</th>
 <th>Title</th>
</tr>
{{range .Items}}
<tr>
    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
    <td>{{.State}}</td>
    <td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
    <td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
</tr>
{{end}}
</table>`

var issueList = template.Must(template.New("issuelist").Parse(templ))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		issueReport(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func issueReport(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var input []string
	for k, v := range r.Form {
		input = append(input, fmt.Sprintf("%s:%v"), k, v[0])
	}
	result, err := github.SearchIssues(input)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
