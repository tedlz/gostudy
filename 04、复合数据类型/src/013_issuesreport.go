package main

import (
	"gostudy/04、复合数据类型/files"
	"log"
	"os"
	"text/template"
	"time"
)

// 013、文本和 HTML 模板 1
// go run 013_issuesreport.go repo:golang/go is:open json decoder override
// 6 issues:
// ---------------------------------------
// Number: 5901
// User:   rsc
// Title:  proposal: encoding/json: allow override type marshaling
// Age:    2360 days
// ---------------------------------------
// Number: 36353
// User:   dsnet
// Title:  proposal: encoding/gob: allow override type marshaling
// Age:    1 days
// ---------------------------------------
// Number: 33835
// User:   Qhesz
// Title:  encoding/json: unmarshalling null into non-nullable golang types
// Age:    129 days
// ---------------------------------------
// Number: 21092
// User:   trotha01
// Title:  encoding/json: unmarshal into slice reuses element data between
// Age:    897 days
// ---------------------------------------
// Number: 5819
// User:   gopherbot
// Title:  encoding/gob: encoder should ignore embedded structs with no exp
// Age:    2377 days
// ---------------------------------------
// Number: 19109
// User:   bradfitz
// Title:  proposal: cmd/go: make fuzzing a first class citizen, like tests
// Age:    1051 days
func main() {
	result, err := files.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// 模板
const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo }} days
{{end}}`

// 时间转换
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// New 用来新建并返回一个模板
// Funcs 用来把自定义函数注册到模板中并返回一个模板
// Parse 用来解析模板
// Must 可以简化错误处理
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))
