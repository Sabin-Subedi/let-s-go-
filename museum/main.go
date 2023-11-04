package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hanldeTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, "Template")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	server.HandleFunc("/template", hanldeTemplate)
	fs := http.FileServer(http.Dir("public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8080", server)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening on port 8080")
}
