package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	err := http.ListenAndServe(":8080", server)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening on port 8080")
}
