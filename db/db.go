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
	Origin_id   int    `bd:"origin_id"`
	EncodedList string `bd:"encoded_list"`
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

	accounts_result := []AccountsDB{}

	err = db.Select(&accounts_result, "SELECT * FROM accounts")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.NamedExec(
		`INSERT INTO accounts (origin_id,account_number,status) 
		VALUES(:origin_id,:account_number,:status)`,
		accounts_result)

	if err != nil {
		log.Fatalln(err)
	}

}
