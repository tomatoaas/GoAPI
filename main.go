package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	openingdao "github.com/tomatoaas/GoAPI/dao/user"
	alldao "github.com/tomatoaas/GoAPI/dao/money"
)
type USER struct {
        USERID         string  `json:"userid"`
        USERNAME       string  `json:"username"`
        PASSWORD        string  `json:"password"`
	WithdrawMoney struct {
		One_yen                 int     `json:"one_yen"`
	        Five_yen                int     `json:"five_yen"`
	        Ten_yen                 int     `json:"ten_yen"`
	        Fifty_yen               int     `json:"fifty_yen"`
	        Hundred_yen             int     `json:"hundred_yen"`
	        Five_hundred_yen        int     `json:"five_hundred_yen"`
        }
}

func main(){
	//ルータのイニシャライズ
	r := mux.NewRouter()

	//ルート（エンドポイント）
	r.HandleFunc("/api/user/", showOpeningIndex).Methods("GET")
	r.HandleFunc("/api/user/add/", adduser).Methods("POST")
	r.HandleFunc("/api/user/update/", updateuser).Methods("POST")
	r.HandleFunc("/api/user/login/", login).Methods("POST")
	r.HandleFunc("/api/money/show/", showMoney).Methods("POST")
	r.HandleFunc("/api/money/add/", addSaving).Methods("POST")

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

func showMoney(w http.ResponseWriter, r *http.Request) {
	var user USER
	json.NewDecoder(r.Body).Decode(&user)

	opening := alldao.ShowAll(user.USERID)
	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}

func addSaving(w http.ResponseWriter, r *http.Request) {
	var user USER
	json.NewDecoder(r.Body).Decode(&user)

	opening := alldao.AddSaving(user.USERID, user.WithdrawMoney.One_yen, user.WithdrawMoney.Five_yen, user.WithdrawMoney.Ten_yen, user.WithdrawMoney.Fifty_yen, user.WithdrawMoney.Hundred_yen, user.WithdrawMoney.Five_hundred_yen)

	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}
