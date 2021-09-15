package dao


import(
	"github.com/tomatoaas/GoAPI/db"
)

//あとでjson形式にするので、jsonタグをあらかじめつけておく

type Add struct {
	Data	string  `json:"Data"`
	Message string  `json:"message"`
}


//ユーザー追加
func AddSaving(userid string, one int, five int, ten int, fif int, hun int, five_hun int) Add {
	db := db.Connect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO SAVING_HISTORY(USER_ID, one_yen, five_yen, ten_yen, fifty_yen, hundred_yen, five_hundred_yen) VALUES(?,?,?,?,?,?,?)", userid, one, five, ten, fif, hun, five_hun)

	var add = Add{}
	if err != nil{
		add.Message = "NO"
	}else {
	        add.Message = "YES"
	}

	return add
}

