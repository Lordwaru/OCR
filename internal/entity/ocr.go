package entity

type AccountsByOriginId struct {
	Origin_id   int            `json:"id"`
	Origin_data string         `json:"encoded_data"`
	Accounts    []AccountsJSON `json:"accounts"`
}

type AccountsJSON struct {
	AccountNumber string `json:"account_number" db:"account_number"`
	Status        string `json:"status" db:"status"`
}

type AccountsDB struct {
	Id             int    `bd:"id"`
	Origin_id      int    `bd:"origin_id"`
	Account_number string `bd:"account_number"`
	Status         string `bd:"status"`
}

type OriginData struct {
	Origin_id    int    `bd:"origin_id"`
	Encoded_list string `bd:"encoded_list"`
}

type EncodedOCR struct {
	Content string `json:"content" binding:"required"`
}
