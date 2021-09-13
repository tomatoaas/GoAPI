package dao

import (
	"github.com/tomatoaas/GoAPI/db"
)

//あとでjson形式にするので、jsonタグをあらかじめつけておく
type Login struct {
	Data struct{
		Username	string  `json:"username"`
	}
	Message string  `json:"message"`
}
//ユーザー更新
func LoginUser(userid string, password string) Login {
	db := db.Connect()
	defer db.Close()

	var log = Login{}
	var pass string
	var username string
	err := db.QueryRow("SELECT USER_NAME, PASSWORD FROM USER WHERE USER_ID = ?",userid).Scan(&username, &pass)
	if err != nil{
		//log.Message = "Not User"
		log.Message = "このユーザーは登録されていません。"
	}else {
		if pass == password{
			log.Data.Username = username;
			log.Message = "OK"
		}else{
			log.Message = "Not Passs"
			//log.Message = "パスワードが間違ってます。"
		}
	}
	return log
}

