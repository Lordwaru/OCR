package ocr

import "fmt"

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
		fmt.Print(r)
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
