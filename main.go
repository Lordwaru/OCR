package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Lordwaru/OCR/accounts"
	"github.com/Lordwaru/OCR/ocr"
	"github.com/Lordwaru/OCR/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//To create a different test file
	//CreateDefaultInputFile("data/default.txt")

	//For files
	//account_list := GetData("data/default.txt")
	//PrintReport(account_list)

	/*
		router := mux.NewRouter().StrictSlash(true)

		router.HandleFunc("/ocr", routes.OcrService)

		handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token", "Authorization", "Access-Control-Allow-Origin"}),
			handlers.ExposedHeaders([]string{""}),
			handlers.MaxAge(1000),
			handlers.AllowCredentials(),
		)(router))

		handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

		http.ListenAndServe(":8080", handler)
		fmt.Println("Service started on port 8080")*/

	router := gin.Default()

	router.POST("/ocr", routes.OcrService)

	router.Run(":8080")

}

func CreateInputFile(amount int, filename string) {
	var ocr_num [9]int
	str := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < amount; i++ {
		for j := range ocr_num {
			ocr_num[j] = rand.Intn(11)
		}
		str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	}

	output := []byte(str)
	err := os.WriteFile(filename, output, 0644)
	check(err)
}

func CreateDefaultInputFile(filename string) {
	var ocr_num [9]int

	str := ""

	ocr_num = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{2, 2, 2, 2, 2, 2, 2, 2, 2}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{3, 3, 3, 3, 3, 3, 3, 3, 3}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{4, 4, 4, 4, 4, 4, 4, 4, 4}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{5, 5, 5, 5, 5, 5, 5, 5, 5}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{6, 6, 6, 6, 6, 6, 6, 6, 6}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{7, 7, 7, 7, 7, 7, 7, 7, 7}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{8, 8, 8, 8, 8, 8, 8, 8, 8}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{9, 9, 9, 9, 9, 9, 9, 9, 9}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{0, 0, 0, 0, 0, 0, 0, 5, 1}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{4, 9, 0, 0, 6, 7, 7, 1, 11}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	ocr_num = [9]int{1, 2, 3, 4, 11, 6, 7, 8, 11}
	str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "

	output := []byte(str)
	err := os.WriteFile(filename, output, 0644)
	check(err)
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
		fmt.Println("Invalid file lenght cannot parse")
		os.Exit(0)
	}

	parsed := make([]accounts.Account, count)

	for i := 0; i < count; i++ {

		ocr_num := ocr.Read(str[i*162 : i*162+162])
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
