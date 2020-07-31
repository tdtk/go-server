package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"encoding/json"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/tdtk/go-server/model"
	"github.com/tdtk/go-server/repository"
)

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params model.LoginFormParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	var pass, userID = repo.GetPasswordByID(params.LoginID)
	if pass == params.Password {

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = params.LoginID
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
		if err != nil {
			http.Error(w, "This password is wrong", 500)
			panic(err)
		}

		var encoder = json.NewEncoder(w)
		m := make(map[string]string)
		m["accessToken"] = tokenString
		m["userID"] = userID
		encoder.Encode(m)
	} else {
		http.Error(w, "This password is wrong!", 500)
	}
	defer repo.Close()
}

func check(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			_, err := token.Method.(*jwt.SigningMethodHMAC)
			if !err {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return []byte(os.Getenv("SIGNINGKEY")), nil
			}
		})
		if err == nil && token.Valid {
			f(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "invalid accessToken", 555)
		}
	}
}

func searchUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params model.SearchFormParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	var users = repo.SearchUser(params)
	var encoder = json.NewEncoder(w)
	encoder.Encode(users)
	defer repo.Close()
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	var user = repo.GetUserByID(userID)
	var encoder = json.NewEncoder(w)
	encoder.Encode(user)
	defer repo.Close()
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.UserInfo
	err := decoder.Decode(&user)
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	repo.UpdateUser(user)
	defer repo.Close()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	repo.DeleteUser(userID)
	defer repo.Close()
}

func getRoleByID(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}
	var repo = repository.NewUserRepository()
	var role = repo.GetRoleByID(roleID)
	var encoder = json.NewEncoder(w)
	encoder.Encode(role)
	defer repo.Close()
}

func getAllRole(w http.ResponseWriter, r *http.Request) {
	var repo = repository.NewUserRepository()
	var roles = repo.GetAllRole()
	var encoder = json.NewEncoder(w)
	encoder.Encode(roles)
	defer repo.Close()
}

func main() {
	http.HandleFunc("/api/login", login)
	http.HandleFunc("/api/search/user", check(searchUser))
	http.HandleFunc("/api/get/user", check(getUserByID))
	http.HandleFunc("/api/update/user", check(updateUser))
	http.HandleFunc("/api/delete/user", check(deleteUser))
	http.HandleFunc("/api/get/role", check(getRoleByID))
	http.HandleFunc("/api/get/all/role", check(getAllRole))
	http.ListenAndServe(":8080", nil)
}
