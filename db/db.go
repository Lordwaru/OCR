package db

import (
	"log"

	"github.com/Lordwaru/OCR/accounts"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AccountsDB struct {
	Id             int    `bd:"id"`
	Origin_id      int    `bd:"origin_id"`
	Account_number string `bd:"account_number"`
	Status         string `bd:"status"`
}

type OriginData struct {
	Origin_id    int    `bd:"origin_id"`
	Encoded_list string `bd:"encoded_list"`
}

/*
Create a rows in DB after receiving data
*/
func Create(encoded_list string, processed_json []accounts.AccountsJSON) {
	//fmt.Println(encoded_list)

	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}

	od := OriginData{}

	od.Encoded_list = encoded_list

	res, err := db.NamedExec(
		`INSERT INTO origin_data (encoded_list) 
		VALUES(:encoded_list)`,
		od)

	if err != nil {
		log.Fatalln(err)
	}

	enclist_id, err := (res.LastInsertId())
	if err != nil {
		log.Fatalln(err)
	}

	adb := []AccountsDB{}

	for _, v := range processed_json {
		adb = append(adb, AccountsDB{
			Origin_id:      int(enclist_id),
			Account_number: v.AccountNumber,
			Status:         v.Status,
		})

	}

	_, err = db.NamedExec(
		`INSERT INTO accounts (origin_id,account_number,status)
	 	VALUES(:origin_id,:account_number,:status)`,
		adb)

	if err != nil {
		log.Fatalln(err)
	}

}
