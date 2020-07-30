package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/tdtk/go-server/model"
	"github.com/tdtk/go-server/repository"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var repo = repository.NewUserRepository()
	var users = repo.FindAllUser()
	for _, user := range users {
		fmt.Fprintf(w, "%+v\n", user)
	}
	fmt.Fprint(w, "Hello from go")
	defer repo.Close()
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params model.LoginFormParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	var pass = repo.GetPasswordByID(params.UserID)
	if pass == params.Password {
		w.Header().Add("token", fmt.Sprintf("%s.%s", params.UserID, params.Password))
	} else {
		http.Error(w, "This password is wrong!", 500)
	}
	defer repo.Close()
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/login", login)
	http.ListenAndServe(":8080", nil)
}
