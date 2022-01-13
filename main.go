package main

import (
	"fmt"
	"os"

	"github.com/Lordwaru/OCR/ocr"
)

func main() {
	var number_a ocr.OCR_number
	var number_b ocr.OCR_number

	number_a.Number = make([]ocr.OCR_pattern, 3)

	number_b = ocr.Zero()

	data, err := os.ReadFile("input.txt")
	check(err)

	fmt.Println((len(data)))
	str := string(data)

	i := 0
	for _, d := range str[0:3] {
		number_a.Number[0].Cells.Characters[i] = d
		i++
	}

	i = 0

	for _, d := range str[27:30] {
		number_a.Number[1].Cells.Characters[i] = d
		i++
	}

	i = 0
	for _, d := range str[54:57] {
		number_a.Number[2].Cells.Characters[i] = d
		i++
	}

	fmt.Println(number_a.Number)
	fmt.Println(number_b.Number)

	if ocr.Compare(number_a, number_b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
