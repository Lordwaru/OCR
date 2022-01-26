package routes

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lordwaru/OCR/accounts"
	"github.com/Lordwaru/OCR/db"

	"github.com/Lordwaru/OCR/ocr"
	"github.com/gin-gonic/gin"
)

type EncodedOCR struct {
	Content string `json:"content" binding:"required"`
}

type Response struct {
	Status  int
	Message string
	Data    interface{}
}

func OcrService(c *gin.Context) {
	var encoded_json EncodedOCR

	if err := c.BindJSON(&encoded_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := GetDataFromEncodedString(encoded_json.Content)

	switch response.Status {
	case http.StatusOK:
		c.JSON(http.StatusOK, response)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Message})
	}

}

func GetDataFromString(str string) Response {

	count, flag := ocr.Count(str)

	if !flag {
		var response Response
		response.Message = "Invalid string length cannot parse"
		response.Data = nil
		response.Status = 500

		return response
	}

	list := make([]accounts.Account, count)

	for i := 0; i < count; i++ {
		ocr_num := ocr.Read(str[i*85 : i*85+83])
		list[i].Number = ocr.ParseToIntArray(ocr_num)
	}

	var json_list []accounts.AccountsJSON

	err_flag := false
	for _, v := range list {
		for _, u := range v.Number {
			if u == 11 {
				err_flag = true
			}
		}

		if !err_flag {
			if accounts.Validate(v) {
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}

				json_list = append(json_list, accounts.AccountsJSON{sb.String(), "OK"})

			} else {
				//664371495 ERR
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				json_list = append(json_list, accounts.AccountsJSON{sb.String(), "ERR"})

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

			json_list = append(json_list, accounts.AccountsJSON{sb.String(), "ILL"})

		}
		err_flag = false

	}

	var response Response
	response.Message = "Success"
	response.Data = json_list
	response.Status = 200

	return response
}

func GetDataFromEncodedString(encoded_list string) Response {

	decoded, err := base64.StdEncoding.DecodeString(encoded_list)

	if err != nil {
		var response Response

		response.Status = 500
		response.Message = "Invalid json object"
		return response
	}

	fmt.Println(len(decoded))
	fmt.Println(decoded)

	count, flag := ocr.CountByte(decoded)

	if !flag {
		var response Response
		response.Status = 500
		response.Message = "Invalid string length"

		return response
	}

	list := make([]accounts.Account, count)

	runes := string(decoded)

	for i := 0; i < count; i++ {

		ocr_num := ocr.Read(runes[i*85 : i*85+83])
		list[i].Number = ocr.ParseToIntArray(ocr_num)
	}

	var json_list []accounts.AccountsJSON

	err_flag := false
	for _, v := range list {
		for _, u := range v.Number {
			if u == 11 {
				err_flag = true
			}
		}

		if !err_flag {
			if accounts.Validate(v) {
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}

				json_list = append(json_list, accounts.AccountsJSON{sb.String(), "OK"})

			} else {
				//664371495 ERR
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				json_list = append(json_list, accounts.AccountsJSON{sb.String(), "ERR"})

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

			json_list = append(json_list, accounts.AccountsJSON{sb.String(), "ILL"})

		}
		err_flag = false

	}

	db.Create(encoded_list, json_list)

	var response Response
	response.Message = "Success"
	response.Data = json_list
	response.Status = 200

	return response
}
