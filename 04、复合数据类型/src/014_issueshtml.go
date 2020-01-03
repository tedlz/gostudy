package main

import (
	"gostudy/04、复合数据类型/files"
	"log"
	"os"
	"text/template"
)

// 014、文本和 HTML 模板 2
// go run 014_issueshtml.go repo:golang/go commenter:gopherbot json encoder >../files/issue.html
// 输出：
// 向 files 文件夹输出一个 issue.html
func main() {
	result, err := files.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
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
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
</table>
`))
