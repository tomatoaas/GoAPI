package dao

import (
	"github.com/tomatoaas/GoAPI/db"
)

//あとでjson形式にするので、jsonタグをあらかじめつけておく
type Update struct {
	Data string  `json:"Data"`
	Message string  `json:"message"`
}
//ユーザー更新
func UpdateUser(userid string, username string) Update {
	db := db.Connect()
	defer db.Close()

	_, err := db.Exec("UPDATE USER set USER_NAME = ? where USER_ID = ?", username, userid)
	var upd = Update{}
	if err != nil{
		upd.Message = err.Error()
	}else {
		upd.Message = "OK"
	}
	return upd
}

