package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Lordwaru/OCR/ocr"
)

func main() {

	//data, err := os.ReadFile("input.txt")
	//check(err)

	//str := string(data)

	/*
		ocr_num := ocr.Read(str)

		parsed := ocr.Parse(ocr_num)

		fmt.Println(parsed)
	*/
	CreateInputFile(9)
}

func CreateInputFile(amount int) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < amount; i++ {
		var ocr_num [9]int

		for i := range ocr_num {

			ocr_num[i] = rand.Intn(10)
		}
		fmt.Println(ocr_num)
		fmt.Println(ocr.IntArrayToString(ocr_num[:]))

	}

	//
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
