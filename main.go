package main

import (
	"github.com/gollira/mux"
	"log"
)

type User struct {
	USER_ID		string	`json:"id"`
	NAME_ID		string	`json:"name"`
	PASSWORD	string	`json:"password"`
}

func main(){
	//ルータのイニシャライズ
	r := mix.NewRouter()

	//ルート（エンドポイント）
	r.HandleFunc("api/user", getUser).Methods("GET")
	r.HandleFunc("api/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
