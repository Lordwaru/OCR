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
)

func main() {
	//CreateInputFile(500, "ocr.txt")

	account_list := GetData("ocr.txt")
	PrintReport(account_list)

}

func CreateInputFile(amount int, filename string) {
	var ocr_num [9]int
	str := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < amount; i++ {
		for j := range ocr_num {
			ocr_num[j] = rand.Intn(10)
		}
		str += ocr.IntArrayToString(ocr_num[:]) + "                                                                                 "
	}

	output := []byte(str)
	err := os.WriteFile(filename, output, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetData(filepath string) []accounts.Account {

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
	err := os.WriteFile("out.txt", output, 0644)
	check(err)
}
