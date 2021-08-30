package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"fmt"

	openingdao "github.com/tomatoaas/GoAPI/dao"
)
type USER struct {
        USER_ID         string  `json:"id"`
        USER_NAME       string  `json:"user_name"`
        PASSWORD        string  `json:"password"`
}

func main(){
	//ルータのイニシャライズ
	r := mux.NewRouter()

	//ルート（エンドポイント）
	r.HandleFunc("/api/user/", showOpeningIndex).Methods("GET")
	r.HandleFunc("/api/user/", adduser).Methods("POST")



	log.Fatal(http.ListenAndServe(":8000", r))
}

func showOpeningIndex(w http.ResponseWriter, r *http.Request) {
	opening := openingdao.GetUsers()
	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(bytes)))
}
func adduser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user USER
	json.NewDecoder(r.Body).Decode(&user)
	opening := openingdao.AddUser(user.USER_NAME, user.PASSWORD)

	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
	//fmt.Fprintf(w, "%s", opening)
}

func updateuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user USER
	json.NewDecoder(r.Body).Decode(&user)
	opening := openingdao.UpdateUser(user.USER_ID, user.USER_NAME)

	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
	//fmt.Fprintf(w, "%s", opening)
}
