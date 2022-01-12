package main

import (
	"fmt"

	"github.com/Lordwaru/OCR/ocr"
)

func main() {
	var number ocr.OCR_number

	number = ocr.One()
	fmt.Println(number)
}
