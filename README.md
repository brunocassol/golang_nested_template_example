# Example of nested template in Go

main.go:

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	})
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "blog", nil)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
```

template.html:

```html
{{define "header"}}
<html>
<head>
    <title>Template example</title>
</head>
<body>
{{end}}

{{define "footer"}}
</body>
</html>
{{end}}

{{define "index"}}
{{template "header"}}
<h1>Index</h1>
<p>Home page. <a href="/blog">Go to blog</a>.</p>
{{template "footer"}}
{{end}}

{{define "blog"}}
{{template "header"}}
<h1>Blog</h1>
<p><a href="/">Go to home page</a>.</p>
<ul>
    <li>Post 3</li>
    <li>Post 2</li>
    <li>Post 1</li>
</ul>
{{template "footer"}}
{{end}}
```
