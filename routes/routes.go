package routes

import (
	"encoding/base64"
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

func CreateAccounts(c *gin.Context) {
	var encoded_json EncodedOCR

	if err := c.BindJSON(&encoded_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := ParseDataFromEncodedString(encoded_json.Content)

	og_id := db.InsertOriginData(encoded_json.Content)
	err := db.InsertAccounts(og_id, response.Data.([]accounts.AccountsJSON))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
	}

	//db.Create(encoded_list, json_list)

	switch response.Status {
	case http.StatusOK:
		c.JSON(http.StatusOK, response)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Message})
	}

}

func GetAccounts(c *gin.Context) {
	accs := db.GetAllAccounts()

	c.JSON(http.StatusOK, accs)
}

func GetAccountsById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resource not valid"})
		return
	}

	ogd := db.GetOriginalDataById(id)
	accs := db.GetAccountsByOriginId(id)

	if false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "resource not found"})
	} else {
		c.JSON(http.StatusOK, accounts.AccountsByOriginId{
			Origin_id:   id,
			Origin_data: ogd,
			Accounts:    accs,
		})
	}

}

func ParseDataFromString(str string) Response {

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

func ParseDataFromEncodedString(encoded_list string) Response {

	decoded, err := base64.StdEncoding.DecodeString(encoded_list)

	if err != nil {
		var response Response

		response.Status = http.StatusBadRequest
		response.Message = "Invalid json object"
		return response
	}

	count, flag := ocr.CountByte(decoded)

	if !flag {
		var response Response
		response.Status = http.StatusBadRequest
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

	var response Response
	response.Message = "Success"
	response.Data = json_list
	response.Status = http.StatusOK

	return response
}
