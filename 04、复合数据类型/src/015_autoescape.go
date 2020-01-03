package main

import (
	"html/template"
	"log"
	"os"
)

// 015、autoescape
// go run 015_autoescape.go >../files/autoescape.html
// 输出：
// 向 files 文件夹输出一个 autoescape.html
func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
