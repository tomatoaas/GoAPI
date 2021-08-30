package dao


import(
	"github.com/tomatoaas/GoAPI/db"
	"github.com/rs/xid"
)

//あとでjson形式にするので、jsonタグをあらかじめつけておく

type Add struct {
	Data 	struct{
		Userid	string  `json:"userid"`
	}
	Message string  `json:"message"`
}

type Data struct {
	Userid	string  `json:"userid"`
}

//ユーザー追加
func AddUser(username string, password string) Add {
	db := db.Connect()
	defer db.Close()
	guid := xid.New()
	uuid := guid.String()

	_, err := db.Exec("INSERT INTO USER VALUES(?,?,?)", uuid, username, password)

	var add = Add{}
	if err != nil{
		add.Message = "No"
	}else {
	        add.Message = "OK"
	}

	add.Data.Userid = uuid

	return add
}

