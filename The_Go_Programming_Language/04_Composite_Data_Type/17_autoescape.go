package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello A!</b>"
	data.B = "<b>Hello B!</b>"
	var report = template.Must(template.New("hello").Parse(templ))
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
