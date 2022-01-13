package ocr

import (
	"strings"
)

type OCR_number struct {
	Number []Row
}

type Row struct {
	Cells struct {
		Characters [3]rune
	}
}

/*Each line of OCR numbers must consist of 162 characters*/
func Count(str string) (int, bool) {
	if len(str)%162 == 0 {
		n := len(str) / 162
		return n, true
	} else {
		return 0, false
	}

}

/* Must be a string of 162 characters */
func Read(str string) []OCR_number {

	ocr_num := make([]OCR_number, 9)

	for i := range ocr_num {
		ocr_num[i].Number = make([]Row, 3)
	}

	n := 0
	x := 0
	y := 0

	for _, r := range str[0:81] {

		ocr_num[n].Number[x].Cells.Characters[y] = r
		y++
		if y > 2 {
			y = 0
			n++
		}
		if n >= 9 {
			n = 0
			x++
		}

		if x > 2 {
			x = 0
		}
	}

	return ocr_num[:]
}

func IntArrayToString(arr []int) string {
	ocr_num := make([]OCR_number, 9)

	for i := range arr {
		switch arr[i] {
		case 0:
			ocr_num[i] = Zero()
		case 1:
			ocr_num[i] = One()
		case 2:
			ocr_num[i] = Two()
		case 3:
			ocr_num[i] = Three()
		case 4:
			ocr_num[i] = Four()
		case 5:
			ocr_num[i] = Five()
		case 6:
			ocr_num[i] = Six()
		case 7:
			ocr_num[i] = Seven()
		case 8:
			ocr_num[i] = Eight()
		case 9:
			ocr_num[i] = Nine()
		default:
			ocr_num[i] = Fake()
		}
	}

	var sb strings.Builder

	for x := 0; x <= 2; x++ {
		for n := 0; n < 9; n++ {
			for y := 0; y <= 2; y++ {
				v := ocr_num[n].Number[x].Cells.Characters[y]
				sb.WriteRune(v)
			}
		}
	}
	/*
		for _, t := range ocr_num {
			for _, u := range t.Number {
				for _, v := range u.Cells.Characters {
					sb.WriteRune(v)
				}
			}
		}
	*/

	return sb.String()
}

func ParseToIntArray(ocr_number []OCR_number) []int {

	var result [9]int

	for i, n := range ocr_number {
		switch {
		case Compare(n, Zero()):
			result[i] = 0
		case Compare(n, One()):
			result[i] = 1
		case Compare(n, Two()):
			result[i] = 2
		case Compare(n, Three()):
			result[i] = 3
		case Compare(n, Four()):
			result[i] = 4
		case Compare(n, Five()):
			result[i] = 5
		case Compare(n, Six()):
			result[i] = 6
		case Compare(n, Seven()):
			result[i] = 7
		case Compare(n, Eight()):
			result[i] = 8
		case Compare(n, Nine()):
			result[i] = 9
		default:
			result[i] = 11
		}

	}

	return result[:]
}

func Compare(A OCR_number, B OCR_number) bool {
	if len(A.Number) > 0 && len(B.Number) > 0 {
		for i, row := range A.Number {
			if row.Cells.Characters[0] != B.Number[i].Cells.Characters[0] {
				return false
			}
			if row.Cells.Characters[1] != B.Number[i].Cells.Characters[1] {
				return false
			}
			if row.Cells.Characters[2] != B.Number[i].Cells.Characters[2] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
