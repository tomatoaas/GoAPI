package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	openingdao "github.com/tomatoaas/GoAPI/dao"
)
type USER struct {
        USERID         string  `json:"userid"`
        USERNAME       string  `json:"username"`
        PASSWORD        string  `json:"password"`
}

func main(){
	//ルータのイニシャライズ
	r := mux.NewRouter()

	//ルート（エンドポイント）
	r.HandleFunc("/api/user/", showOpeningIndex).Methods("GET")
	r.HandleFunc("/api/user/add/", adduser).Methods("POST")
	r.HandleFunc("/api/user/update/", updateuser).Methods("POST")
	r.HandleFunc("/api/user/login/", login).Methods("POST")

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
	opening := openingdao.AddUser(user.USERNAME, user.PASSWORD)

	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}

func updateuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user USER
	json.NewDecoder(r.Body).Decode(&user)
	opening := openingdao.UpdateUser(user.USERID, user.USERNAME)

	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user USER
	json.NewDecoder(r.Body).Decode(&user)
	opening := openingdao.LoginUser(user.USERID, user.PASSWORD)

	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}
