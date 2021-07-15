package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	openingdao "github.com/tomatoaas/GoAPI/dao"
)
type USER struct {
        USER_ID         string  `json:"id"`
        USER_NAME       string  `json:"name"`
        PASSWORD        string  `json:"password"`
}

//Userのデータを保持するスライスの作成
var users []USER
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
	json.NewEncoder(w).Encode(&USER{})
}

//新規ユーザーの追加
func createUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user USER
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
			var user USER
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
	r := mux.NewRouter()

	//ルート（エンドポイント）
	r.HandleFunc("/api/user/", showOpeningIndex).Methods("GET")
	r.HandleFunc("api/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func showOpeningIndex(w http.ResponseWriter, r *http.Request) {
	opening := openingdao.getUsers()
	//json形式に変換します
	bytes, err :=json.Marshal(opening)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(bytes)))
}
