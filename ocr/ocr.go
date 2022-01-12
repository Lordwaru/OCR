package ocr

type OCR_number struct {
	number []OCR_pattern
}

type OCR_pattern struct {
	Cells struct {
		characters [3]rune
	}
}

func Compare(A OCR_number, B OCR_number) bool {
	for i, row := range A.number {
		if row.Cells.characters[0] != B.number[i].Cells.characters[0] {
			return false
		}
		if row.Cells.characters[1] != B.number[i].Cells.characters[1] {
			return false
		}
		if row.Cells.characters[2] != B.number[i].Cells.characters[2] {
			return false
		}
	}

	return true
}
