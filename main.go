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
	var pass = repo.GetPasswordByID(params.LoginID)
	if pass == params.Password {
		var encoder = json.NewEncoder(w)
		m := make(map[string]string)
		m["accessToken"] = fmt.Sprintf("%s.%s", params.LoginID, params.Password)
		encoder.Encode(m)
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
