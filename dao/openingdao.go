package dao

import (
	"github.com/tomatoaas/GoAPI/db"
	"github.com/rs/xid"
	_ "fmt"
)

//User型の構造体
//あとでjson形式にするので、jsonタグをあらかじめつけておく
//type Data struct {
//	userid	string	`json:"userid"`
//}
//type Add struct {
//	Datas Data
//	message		string	`json:"message"`
//}

type Add struct {
	userid		string	`json:"userod"`
	message		string 	`json:"message"`
}

type Opening struct {
        USER_ID         string  `json:"id"`
        USER_NAME       string  `json:"name"`
        PASSWORD        string  `json:"password"`
}

func GetUsers() []Opening {
	db := db.Connect()
	defer db.Close()

	//row を取得
	rows, err := db.Query("SELECT * FROM USER");
	if err != nil {
		panic(err.Error())
	}

	//User型のスライスに格納
	openingArgs := make([]Opening, 0)
	for rows.Next() {
		var opening Opening
		err = rows.Scan(&opening.USER_ID, &opening.USER_NAME, &opening.PASSWORD)
		if err != nil {
			panic(err.Error())
		}
		openingArgs = append(openingArgs, opening)
	}
	return openingArgs
}

func AddUser(username string,pass string)[]Add {
	db := db.Connect()
	defer db.Close()
	guid := xid.New()
	uuid := guid.String()

	_, err := db.Exec("INSERT INTO USER VALUES(?,?,?)",uuid, username, pass)

	if err != nil{
//		message := "NO"
	}else {
//		message := "YES"
	}

//	data := Data{userid: uuid}
	//add := Add{Datas: data, message: "YES"}
	add := Add{uuid, "YES"}

	Adds := make([]Add, 0)

	Adds = append(Adds, add)

	return Adds

}

