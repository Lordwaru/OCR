package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Lordwaru/OCR/accounts"
	"github.com/Lordwaru/OCR/ocr"
)

func main() {
	CreateInputFile(9, "ocr.txt")

	account_list := GetData("ocr.txt")

	for _, v := range account_list {
		fmt.Println(v)
		fmt.Println(accounts.Validate(v))
	}

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
