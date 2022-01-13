package main

import (
	"fmt"
	"os"

	"github.com/Lordwaru/OCR/ocr"
)

func main() {

	data, err := os.ReadFile("input.txt")
	check(err)

	str := string(data)

	ocr_num := ocr.GetDigits(str)

	fmt.Println(ocr_num)

	for _, number := range ocr_num {
		if ocr.Compare(number, ocr.Zero()) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
