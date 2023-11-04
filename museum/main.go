package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening on port 8080")
}
