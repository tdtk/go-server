package main

import (
	"fmt"
	"net/http"

	"github.com/tdtk/go-server/repository"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	repository.Connect()
	fmt.Fprint(w, "Hello from go")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
