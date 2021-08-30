package openingdao


import(
	"github.com/tomatoaas/go_practice/pkg/db"
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

	_, err := db.Exec("INSERT INTO user VALUES(?,?,?)", uuid, username, password)

	var add = Add{}
	if err != nil{
		add.Message = "No"
	}else {
        add.Message = "OK"
	}

	add.Data.Userid = uuid

	return add
}

