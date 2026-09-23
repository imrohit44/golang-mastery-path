package main

import (
	"html/template"
	"os"
)

type PageData struct {
	Title string
	Items []string
}

func main() {
	// Define an HTML template string
	const tpl = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
	<h1>{{.Title}}</h1>
	<ul>
		{{range .Items}}
		<li>{{.}}</li>
		{{else}}
		<li>No items found</li>
		{{end}}
	</ul>
</body>
</html>`

	// Parse the template
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		panic(err)
	}

	// Supply data to the template and output to standard out (or an HTTP response)
	data := PageData{
		Title: "My Go Hardware Projects",
		Items: []string{"Arduino Rover", "ESP32 Sensor", "Raspberry Pi Server"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}