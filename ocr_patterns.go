package main

type OCR_pattern struct {
	Cells struct {
		characters [4]rune
	}
}

type OCR_number struct {
	number []OCR_pattern
}

/*

|
|
*/
func One() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', ' ', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{' ', ' ', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{' ', ' ', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
 _
 _|
|_
*/
func Two() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{' ', '_', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{'|', '_', ' ', ' '}
	o_number.number = cheatsheet[:]

	return o_number
}

/*
_
_|
_|
*/
func Three() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{' ', '_', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{' ', '_', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*

|_|
  |
*/
func Four() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', ' ', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{'|', '_', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{' ', ' ', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
 _
|_
 _|
*/
func Five() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{'|', '_', ' ', ' '}
	cheatsheet[2].Cells.characters = [4]rune{' ', '_', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
 _
|_
|_|
*/
func Six() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{'|', '_', ' ', ' '}
	cheatsheet[2].Cells.characters = [4]rune{'|', '_', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
_
 |
 |
*/
func Seven() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{' ', ' ', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{' ', ' ', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
 _
|_|
|_|
*/

func Eight() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{'|', '_', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{'|', '_', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
 _
|_|
 _|
*/

func Nine() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{'|', '_', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{' ', '_', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}

/*
 _
| |
|_|
*/

func Zero() OCR_number {
	var o_number OCR_number
	var cheatsheet [3]OCR_pattern

	cheatsheet[0].Cells.characters = [4]rune{' ', '_', ' ', ' '}
	cheatsheet[1].Cells.characters = [4]rune{'|', ' ', '|', ' '}
	cheatsheet[2].Cells.characters = [4]rune{'|', '_', '|', ' '}

	o_number.number = cheatsheet[:]

	return o_number
}
