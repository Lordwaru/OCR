package ocr

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lordwaru/OCR/internal/entity"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func PostAccounts(c *gin.Context) {
	var encoded_json entity.EncodedOCR

	if err := c.BindJSON(&encoded_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accounts, err := ParseDataFromEncodedString(encoded_json.Content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	og_id := InsertOriginData(encoded_json.Content)
	err = InsertAccounts(og_id, accounts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	//db.Create(encoded_list, json_list)

	c.JSON(http.StatusOK, accounts)

}

func GetAllAccounts(c *gin.Context) {
	accs := SelectAllAccounts()

	c.JSON(http.StatusOK, accs)
}

func SelectAllAccounts() []entity.AccountsByOriginId {
	result := []entity.AccountsByOriginId{}

	ogd_result := []entity.OriginData{}
	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Select(&ogd_result, "SELECT * FROM origin_data")

	for _, v := range ogd_result {
		accs := GetAccountsByOriginId(v.Origin_id)
		result = append(result, entity.AccountsByOriginId{
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

func SelectAccountsById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resource not valid"})
		return
	}

	ogd := GetOriginalDataById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	accs := GetAccountsByOriginId(id)

	if false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "resource not found"})
	} else {
		c.JSON(http.StatusOK, entity.AccountsByOriginId{
			Origin_id:   id,
			Origin_data: ogd,
			Accounts:    accs,
		})
	}

}

//checksum calculation (1*d1 + 2*d2 + 3*d3 + … + 9*d9) mod 11 = 0
func ValidateAccount(acc entity.Account) bool {
	var checksum int
	for i := 0; i < len(acc.Number); i++ {
		checksum += (9 - i) * acc.Number[i]

	}

	return checksum%11 == 0
}

func IntArrayToString(arr []int) string {
	ocr_num := make([]OCR_number, 9)

	for i := range arr {
		switch arr[i] {
		case 0:
			ocr_num[i] = Zero()
		case 1:
			ocr_num[i] = One()
		case 2:
			ocr_num[i] = Two()
		case 3:
			ocr_num[i] = Three()
		case 4:
			ocr_num[i] = Four()
		case 5:
			ocr_num[i] = Five()
		case 6:
			ocr_num[i] = Six()
		case 7:
			ocr_num[i] = Seven()
		case 8:
			ocr_num[i] = Eight()
		case 9:
			ocr_num[i] = Nine()
		default:
			ocr_num[i] = Fake()
		}
	}

	var sb strings.Builder

	for x := 0; x <= 2; x++ {
		for n := 0; n < 9; n++ {
			for y := 0; y <= 2; y++ {
				v := ocr_num[n].Number[x].Cells.Characters[y]
				sb.WriteRune(v)
			}
			sb.WriteRune('\n')
		}
	}

	return sb.String()
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

	od := entity.OriginData{}

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
func InsertAccounts(og_id int, processed_json []entity.AccountsJSON) error {

	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	adb := []entity.AccountsDB{}

	for _, v := range processed_json {
		adb = append(adb, entity.AccountsDB{
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

func GetOriginalDataById(og_id int) string {
	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}

	row := db.QueryRowx("SELECT * FROM origin_data WHERE origin_id= ?", og_id)

	var ogd entity.OriginData
	err = row.Scan(&ogd.Origin_id, &ogd.Encoded_list)

	if err == sql.ErrNoRows {
		//log.Fatal(err)
	}

	return ogd.Encoded_list

}

func GetAccountsByOriginId(og_id int) []entity.AccountsJSON {
	db, err := sqlx.Connect("mysql", "ocr:qweqwe@(localhost:3306)/ocr")

	if err != nil {
		log.Fatalln(err)
	}
	var accs []entity.AccountsJSON

	err = db.Select(&accs, "SELECT account_number,status FROM accounts WHERE origin_id= ?", og_id)

	return accs
}
