package ocr

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Lordwaru/OCR/internal/entity"
)

type OCR_number struct {
	Number []Row
}

type Row struct {
	Cells struct {
		Characters [3]rune
	}
}
type Response struct {
	Status  int
	Message string
	Data    interface{}
}

/*
Each line of OCR numbers must consist of 85 characters, except the last number that's why
we pad the lenght with +2 to compensate for the missing last two line jump (\n)
*/
func Count(str string) (int, error) {
	if (len(str)+2)%85 == 0 && len(str) != 0 {
		n := (len(str) + 2) / 85
		return n, nil
	} else {
		return 0, errors.New("Invalid lenght")
	}
}

/*
Each line of OCR numbers must consist of 85 characters, except the last number that's why
we pad the lenght with +2 to compensate for the missing last two line jump (\n)
*/
func CountByte(str []byte) (int, bool) {
	if (len(str)+2)%85 == 0 && len(str) != 0 {
		n := (len(str) + 2) / 85
		return n, true
	} else {
		return 0, false
	}

}

/*
Must be a string of 83 characters
*/
func Read(str string) []OCR_number {

	ocr_num := make([]OCR_number, 9)

	for i := range ocr_num {
		ocr_num[i].Number = make([]Row, 3)
	}

	n := 0
	x := 0
	y := 0

	for _, r := range str[:] {

		if r == '\n' {
			continue
		}
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

/*
Parse OCR number to an Integer Array
*/
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

/*
Compare two OCR numbers
*/
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

/*
Parse OCR number to an Accounts Array
*/
func ParseDataFromString(str string) ([]entity.AccountsJSON, error) {

	count, err := Count(str)

	if err != nil {

		return nil, err
	}

	list := make([]entity.Account, count)

	for i := 0; i < count; i++ {
		ocr_num := Read(str[i*85 : i*85+83])
		list[i].Number = ParseToIntArray(ocr_num)
	}

	var json_list []entity.AccountsJSON

	err_flag := false
	for _, v := range list {
		for _, u := range v.Number {
			if u == 11 {
				err_flag = true
			}
		}

		if !err_flag {
			if ValidateAccount(v) {
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}

				json_list = append(json_list, entity.AccountsJSON{sb.String(), "OK"})

			} else {
				//664371495 ERR
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				json_list = append(json_list, entity.AccountsJSON{sb.String(), "ERR"})

			}
		} else {
			//86110??36 ILL
			var sb strings.Builder
			for _, n := range v.Number {
				if n != 11 {
					sb.WriteString(strconv.Itoa(n))
				} else {
					sb.WriteString("?")
				}

			}

			json_list = append(json_list, entity.AccountsJSON{sb.String(), "ILL"})

		}
		err_flag = false

	}

	return json_list, nil
}

/*
Parse OCR number to an Accounts Array
*/
func ParseDataFromEncodedString(encoded_list string) ([]entity.AccountsJSON, error) {

	decoded, err := base64.StdEncoding.DecodeString(encoded_list)

	if err != nil {

		return nil, err
	}

	return ParseDataFromString(string(decoded))

}

func GetDataFromFile(filepath string) []entity.Account {

	data, err := os.ReadFile(filepath)
	check(err)

	str := string(data)

	count, err := Count(str)

	if err != nil {
		fmt.Println("Invalid file length cannot parse")
		os.Exit(0)
	}

	parsed := make([]entity.Account, count)

	for i := 0; i < count; i++ {

		ocr_num := Read(str[i*85 : i*85+83]) //disregard the last two linejumps
		parsed[i].Number = ParseToIntArray(ocr_num)
	}

	return parsed
}

func PrintReport(account_list []entity.Account) {
	var sb strings.Builder
	err_flag := false
	for _, v := range account_list {
		for _, u := range v.Number {
			if u == 11 {
				err_flag = true
			}
		}
		if !err_flag {
			if ValidateAccount(v) {
				//print 457508000 OK
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				sb.WriteString(" ")
				sb.WriteString("OK")
				sb.WriteString("\n")

			} else {
				//664371495 ERR
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				sb.WriteString(" ")
				sb.WriteString("ERR")
				sb.WriteString("\n")

			}
		} else {
			//86110??36 ILL
			for _, n := range v.Number {
				if n != 11 {
					sb.WriteString(strconv.Itoa(n))
				} else {
					sb.WriteString("?")
				}

			}
			sb.WriteString(" ")
			sb.WriteString("ILL")
			sb.WriteString("\n")

		}
		err_flag = false

	}
	output := []byte(sb.String())
	err := os.WriteFile("data/out.txt", output, 0644)
	check(err)
}
