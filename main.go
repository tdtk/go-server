package main

import (
	"fmt"
	"net/http"

	"github.com/tdtk/go-server/repository"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var repo = repository.NewUserRepository()
	var users = repo.FindAllUser()
	for _, user := range users {
		fmt.Fprintf(w, "%+v\n", user)
	}
	fmt.Fprint(w, "Hello from go")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
