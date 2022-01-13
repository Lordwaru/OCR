package accounts

type Account struct {
	Number []int
}

//checksum calculation (1*d1 + 2*d2 + 3*d3 + … + 9*d9) mod 11 = 0
func Validate(acc Account) bool {
	return ((9*acc.Number[0]+
		8*acc.Number[1]+
		7*acc.Number[2]+
		6*acc.Number[3]+
		5*acc.Number[4]+
		4*acc.Number[5]+
		3*acc.Number[6]+
		2*acc.Number[7]+
		1*acc.Number[8])%11 == 0)

}
