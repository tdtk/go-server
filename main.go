package main

import (
	"fmt"
	"github.com/tdtk/go-server/repository"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	connect()
	fmt.Fprint(w, "Hello from go")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
