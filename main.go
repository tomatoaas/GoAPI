package main

import (
	"github.com/gollira/mux"
	"log"
)

func main(){
	//ルータのイニシャライズ
	r := mix.NewRouter()

	//ルート（エンドポイント）
	r.HandleFunc("api/user", getUser).Methods("GET")
	r.HandleFunc("api/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
