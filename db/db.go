package db

import (
	"log"

	"github.com/Lordwaru/OCR/accounts"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type OriginData struct {
	EncodedList string `bd:"encoded_list"`
}

/*
Create a rows in DB after receiving data
*/
func Create(encoded_list string, processed_json []accounts.AccountsJSON) {
	//fmt.Println(encoded_list)
	_, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}
}
