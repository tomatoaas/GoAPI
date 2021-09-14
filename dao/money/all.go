package dao

import (
	"github.com/tomatoaas/GoAPI/db"
)

type Data struct {
	Data	[]All	`json:"data"`
	Message	string		`json:"message"`
}

type All struct {
        SAVING_DATETIME         string	`json:"datetime"`
        One_yen			int	`json:"one_yen"`
        Five_yen	        int	`json:"five_yen"`
        Ten_yen			int	`json:"ten_yen"`
        Fifty_yen		int	`json:"fifty_yen"`
        Hundred_yen		int	`json:"hundred_yen"`
        Five_hundred_yen        int	`json:"five_hundred_yen"`
	Money			int	`json:"money"`
}

func ShowAll(userid string) Data {
	db := db.Connect()
	defer db.Close()
	var erflg = 0

	//row を取得
	rows, err := db.Query("SELECT SAVING_DATETIME, one_yen, five_yen, ten_yen, fifty_yen, hundred_yen, five_hundred_yen FROM SAVING_HISTORY WHERE USER_ID = ?", userid);
	if err != nil {
		erflg = 1
		panic(err.Error())
	}

	//User型のスライスに格納
	allArgs := make([]All, 0)
	for rows.Next() {
		var all All
		err = rows.Scan(&all.SAVING_DATETIME, &all.One_yen, &all.Five_yen, &all.Ten_yen, &all.Fifty_yen, &all.Hundred_yen, &all.Five_hundred_yen)
		all.Money = (all.One_yen * 1 + all.Five_yen * 5 + all.Ten_yen * 10 + all.Fifty_yen * 50 + all.Hundred_yen * 100 + all.Five_hundred_yen * 500)
		if err != nil {
			erflg = 1
			panic(err.Error())
		}
		allArgs = append(allArgs, all)
	}
	var data = Data{}
	data.Data = allArgs
	if erflg == 1 {
		data.Message = "NO"
	}else{
		data.Message = "YES"
	}
	return data
}


