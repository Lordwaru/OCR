package ocr

/*

	' ' = 32,
	'|' = 124,
	'_' = 95
	'/10'

	000000000
	_  _  _  _  _  _  _  _  _ | || || || || || || || || ||_||_||_||_||_||_||_||_||_|

	100000000

	    _  _  _  _  _  _  _  _   || || || || || || || || |  ||_||_||_||_||_||_||_||_|

	200000000
	 _  _  _  _  _  _  _  _  _ _|| || || || || || || || ||_ |_||_||_||_||_||_||_||_|

	300000000
	 _  _  _  _  _  _  _  _  _
	 _|| || || || || || || || |
	 _||_||_||_||_||_||_||_||_|
	    _  _  _  _  _  _  _  _
	  || || || || || || || || |
	  ||_||_||_||_||_||_||_||_|

    _  _  _  _  _  _  _  _
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
