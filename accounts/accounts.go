package accounts

type Account struct {
	Number []int
}

type AccountsByOriginId struct {
	Origin_id   int            `json:"id"`
	Origin_data string         `json:"encoded_data"`
	Accounts    []AccountsJSON `json:"accounts"`
}

type AccountsJSON struct {
	AccountNumber string `json:"account_number" db:"account_number"`
	Status        string `json:"status" db:"status"`
}

//checksum calculation (1*d1 + 2*d2 + 3*d3 + … + 9*d9) mod 11 = 0
func Validate(acc Account) bool {
	var checksum int
	for i := 0; i < len(acc.Number); i++ {
		checksum += (9 - i) * acc.Number[i]

	}

	return checksum%11 == 0
}
