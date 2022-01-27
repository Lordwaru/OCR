package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Lordwaru/OCR/accounts"
	"github.com/Lordwaru/OCR/ocr"
	"github.com/Lordwaru/OCR/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/ocr/create", routes.CreateAccounts)
	router.GET("/ocr/accounts/", routes.GetAccounts)
	router.GET("/ocr/accounts/:id", routes.GetAccountsById)

	router.Run(":8080")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetDataFromFile(filepath string) []accounts.Account {

	data, err := os.ReadFile(filepath)
	check(err)

	str := string(data)

	count, flag := ocr.Count(str)

	if !flag {
		fmt.Println("Invalid file length cannot parse")
		os.Exit(0)
	}

	parsed := make([]accounts.Account, count)

	for i := 0; i < count; i++ {

		ocr_num := ocr.Read(str[i*85 : i*85+83]) //disregard the last two linejumps
		parsed[i].Number = ocr.ParseToIntArray(ocr_num)
	}

	return parsed
}

func PrintReport(account_list []accounts.Account) {
	var sb strings.Builder
	err_flag := false
	for _, v := range account_list {
		for _, u := range v.Number {
			if u == 11 {
				err_flag = true
			}
		}
		if !err_flag {
			if accounts.Validate(v) {
				//print 457508000 OK
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				sb.WriteString(" ")
				sb.WriteString("OK")
				sb.WriteString("\n")

			} else {
				//664371495 ERR
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				sb.WriteString(" ")
				sb.WriteString("ERR")
				sb.WriteString("\n")

			}
		} else {
			//86110??36 ILL
			for _, n := range v.Number {
				if n != 11 {
					sb.WriteString(strconv.Itoa(n))
				} else {
					sb.WriteString("?")
				}

			}
			sb.WriteString(" ")
			sb.WriteString("ILL")
			sb.WriteString("\n")

		}
		err_flag = false

	}
	output := []byte(sb.String())
	err := os.WriteFile("data/out.txt", output, 0644)
	check(err)
}
