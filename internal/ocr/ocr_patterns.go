package ocr

import (
	"math/rand"
	"os"
	"time"
)

/*
cheatsheet
	' ' = 32,
	'|' = 124,
	'_' = 95
	'/10'
*/

func Fake() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', ' ', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'_', '_', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{' ', ' ', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
|
|
*/
func One() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', ' ', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{' ', ' ', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{' ', ' ', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
 _
 _|
|_
*/
func Two() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{' ', '_', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{'|', '_', ' '}
	o_number.Number = cheatsheet[:]

	return o_number
}

/*
_
_|
_|
*/
func Three() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{' ', '_', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{' ', '_', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*

|_|
  |
*/
func Four() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', ' ', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'|', '_', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{' ', ' ', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
 _
|_
 _|
*/
func Five() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'|', '_', ' '}
	cheatsheet[2].Cells.Characters = [3]rune{' ', '_', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
 _
|_
|_|
*/
func Six() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'|', '_', ' '}
	cheatsheet[2].Cells.Characters = [3]rune{'|', '_', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
_
 |
 |
*/
func Seven() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{' ', ' ', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{' ', ' ', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
 _
|_|
|_|
*/

func Eight() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'|', '_', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{'|', '_', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
 _
|_|
 _|
*/

func Nine() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'|', '_', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{' ', '_', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

/*
 _
| |
|_|
*/

func Zero() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]Row

	cheatsheet[0].Cells.Characters = [3]rune{' ', '_', ' '}
	cheatsheet[1].Cells.Characters = [3]rune{'|', ' ', '|'}
	cheatsheet[2].Cells.Characters = [3]rune{'|', '_', '|'}

	o_number.Number = cheatsheet[:]

	return o_number
}

func CreateInputFile(amount int, filename string) {
	var ocr_num [9]int
	str := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < amount; i++ {
		for j := range ocr_num {
			ocr_num[j] = rand.Intn(11)
		}
		if i < amount {
			str += IntArrayToString(ocr_num[:]) + "\n"
		} else {
			str += IntArrayToString(ocr_num[:])
		}

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

func CreateDefaultInputFile(filename string) {
	var ocr_num [9]int

	str := ""

	for i := 0; i <= 9; i++ {
		ocr_num = [9]int{i, i, i, i, i, i, i, i, i}
	}
	ocr_num = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	str += IntArrayToString(ocr_num[:]) + "\n"
	ocr_num = [9]int{0, 0, 0, 0, 0, 0, 0, 5, 1}
	str += IntArrayToString(ocr_num[:]) + "\n"
	ocr_num = [9]int{4, 9, 0, 0, 6, 7, 7, 1, 11}
	str += IntArrayToString(ocr_num[:]) + "\n"
	ocr_num = [9]int{1, 2, 3, 4, 11, 6, 7, 8, 11}
	str += IntArrayToString(ocr_num[:])

	output := []byte(str)
	err := os.WriteFile(filename, output, 0644)
	check(err)
}
