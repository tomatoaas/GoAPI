package dao

import "github.com/tomatoaas/GoAPI/db"

//User型の構造体
//あとでjson形式にするので、jsonタグをあらかじめつけておく
type Opening struct {
        USER_ID         string  `json:"id"`
        USER_NAME       string  `json:"name"`
        PASSWORD        string  `json:"password"`
}

func getUsers() []Opening {
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

