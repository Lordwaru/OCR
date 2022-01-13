package ocr

type OCR_number struct {
	Number []OCR_pattern
}

type OCR_pattern struct {
	Cells struct {
		Characters [3]rune
	}
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
