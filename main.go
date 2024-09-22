package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Hello World Page",
		Message: "Hello, World",
	}

	tmpl.Execute(w, data)
}

func goodbyeWorldHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Goodbye World Page",
		Message: "Goodbye, World",
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
    http.HandleFunc("/goodbye", goodbyeWorldHandler)

	fmt.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
