package main

import (
	"github.com/gollira/mux"
	"log"
)

type User struct {
	USER_ID		string	`json:"id"`
	USER_NAME	string	`json:"name"`
	PASSWORD	string	`json:"password"`
}

//Userのデータを保持するスライスの作成
var users []User

func main(){
	//ルータのイニシャライズ
	r := mix.NewRouter()

	//モックデータの作成
	users = append(users, User{USER_ID: "1", USER_NAME: "Yoshi", PASSWORD: "123qwe"})
	users = append(users, User{USER_ID: "2", USER_NAME: "mura", PASSWORD: "123qwecc"})


	//ルート（エンドポイント）
	r.HandleFunc("api/user", getUser).Methods("GET")
	r.HandleFunc("api/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
