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

Create a new origin data registry in DB and return its ID
*/
func InsertOriginData(encoded_list string) int {
	//fmt.Println(encoded_list)

	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}

	od := OriginData{}

	od.Encoded_list = encoded_list

	res, err := db.NamedExec(
		`INSERT INTO 
		origin_data (encoded_list) 
		VALUES
		(:encoded_list)`,
		od)

	if err != nil {
		log.Fatalln(err)
	}

	enclist_id, err := (res.LastInsertId())

	if err != nil {
		log.Fatalln(err)
	}

	return int(enclist_id)

}

/*
Create a new  registry in DB and return its ID
*/
func InsertAccounts(og_id int, processed_json []accounts.AccountsJSON) error {

	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	adb := []AccountsDB{}

	for _, v := range processed_json {
		adb = append(adb, AccountsDB{
			Origin_id:      int(og_id),
			Account_number: v.AccountNumber,
			Status:         v.Status,
		})

	}

	_, err = db.NamedExec(
		`INSERT INTO 
		accounts (origin_id,account_number,status)
	 	VALUES
		 (:origin_id,:account_number,:status)`,
		adb)

	if err != nil {
		return err
	} else {
		return nil
	}

}

func GetAllAccounts() []accounts.AccountsByOriginId {
	result := []accounts.AccountsByOriginId{}

	ogd_result := []OriginData{}
	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Select(&ogd_result, "SELECT * FROM origin_data")

	for _, v := range ogd_result {
		accs := GetAccountsByOriginId(v.Origin_id)
		result = append(result, accounts.AccountsByOriginId{
			Origin_id:   v.Origin_id,
			Origin_data: v.Encoded_list,
			Accounts:    accs,
		})

	}

	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func GetOriginalDataById(og_id int) string {
	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}
	var ogd OriginData

	err = db.Select(&ogd, "SELECT * FROM accounts WHERE origin_id= ?", og_id)

	return ogd.Encoded_list

}

func GetAccountsByOriginId(og_id int) []accounts.AccountsJSON {
	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}
	var accs []accounts.AccountsJSON

	err = db.Select(&accs, "SELECT account_number,status FROM accounts WHERE origin_id= ?", og_id)

	return accs
}
