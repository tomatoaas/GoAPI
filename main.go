package main

import (
	"encoding/json"
	"github.com/gollira/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	USER_ID		string	`json:"id"`
	USER_NAME	string	`json:"name"`
	PASSWORD	string	`json:"password"`
}

//Userのデータを保持するスライスの作成
var users []User

//すべてのuserを取得する
func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

//特定のuserを取得する

func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//次のIDがある限りループ
	for _, item := range users{
		if item.USER_ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

//新規ユーザーの追加
func createUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user User
	_= json.NewDecoder(r.Body).Decode(&user)
	user.USER_ID = strconv.Itoa(rand.Intn(10000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

//ユーザーデータの更新
func updateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range users{
		if item.USER_ID == params["id"]{
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.USER_ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

//ユーザーの削除
func deleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range users{
		if item.USER_ID == params["id"]{
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}


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
