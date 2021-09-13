package dao

import (
	"github.com/tomatoaas/GoAPI/db"
)

type All struct {
        SAVING_DATETIME         string  `json:"time"`
        one_yen       string  `json:"one"`
        five_yen        string  `json:"five"`
        ten_yen        string  `json:"ten"`
        fifty_yen        string  `json:"fif"`
        hundred_yen        string  `json:"hun"`
        five_hundred_yen        string  `json:"five_hun"`
}

func ShowAll(userid string) []All {
	db := db.Connect()
	defer db.Close()

	//row を取得
	rows, err := db.Query("SELECT SAVING_DATETIME, one_yen, five_yen, ten_yen, fifty_yen, hundred_yen, five_hundred_yen FROM SAVING_HISTORY WHERE USER_ID = ?", userid);
	if err != nil {
		panic(err.Error())
	}

	//User型のスライスに格納
	allArgs := make([]All, 0)
	for rows.Next() {
		var all All
		err = rows.Scan(&all.SAVING_DATETIME, &all.one_yen, &all.five_yen, &all.ten_yen, &all.fifty_yen, &all.hundred_yen, &all.five_hundred_yen)
		if err != nil {
			panic(err.Error())
		}
		allArgs = append(allArgs, all)
	}
	return allArgs
}


