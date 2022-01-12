package main

import (
	"fmt"

	"github.com/Lordwaru/OCR/ocr"
)

func main() {
	var number_a ocr.OCR_number
	var number_b ocr.OCR_number

	number_a = ocr.One()
	number_b = ocr.Two()

	if ocr.Compare(number_a, number_b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
