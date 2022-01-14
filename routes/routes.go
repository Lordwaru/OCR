package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lordwaru/OCR/accounts"
	"github.com/Lordwaru/OCR/ocr"
)

type AccountsJSON struct {
	AccountNumber string `json:"account_number"`
	Status        string `json:"status"`
}

type Response struct {
	Status  int
	Message string
	Data    interface{}
}

func OcrService(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	response := GetDataFromString(string(body))

	switch response.Status {
	case http.StatusOK:
		w.WriteHeader(http.StatusOK)
		json_str, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(json_str))
	default:
		json_str, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			w.WriteHeader(response.Status)
			fmt.Fprintln(w, string(json_str))
		}
	}

	if err != nil {
		log.Panic(err)
	}

}

func GetDataFromString(str string) Response {

	count, flag := ocr.Count(str)

	if !flag {
		var response Response
		response.Message = "Invalid file length cannot parse"
		response.Data = nil
		response.Status = 500

		return response
	}

	list := make([]accounts.Account, count)

	for i := 0; i < count; i++ {

		ocr_num := ocr.Read(str[i*162 : i*162+162])
		list[i].Number = ocr.ParseToIntArray(ocr_num)
	}

	var json_list []AccountsJSON

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

				json_list = append(json_list, AccountsJSON{sb.String(), "0K"})

			} else {
				//664371495 ERR
				var sb strings.Builder
				for _, n := range v.Number {
					sb.WriteString(strconv.Itoa(n))
				}
				json_list = append(json_list, AccountsJSON{sb.String(), "ERR"})

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

			json_list = append(json_list, AccountsJSON{sb.String(), "ILL"})

		}
		err_flag = false

	}

	var response Response
	response.Message = "Success"
	response.Data = json_list
	response.Status = 200

	return response
}
