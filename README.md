OCR Reader

Reads numbers in a similar to a digital digits display


Recognizes numbers following a pattern that's defined on ocr_patterns, and stores them in an integer array 
if the pattern doesn't match any number it stores an double digit number(11).

Evaluates a checksum like this

account number: 3  4  5  8  8  2  8  6  5
position name: d9 d8 d7 d6 d5 d4 d3 d2 d1

checksum calculation:
(1*d1 + 2*d2 + 3*d3 + … + 9*d9) mod 11 == 0

and returns 

999999999 ERR If it doesnt fullfil the condition
123456789 OK  If it fullfils the condition

If there was a non reconized number patter it returns

49006771? ILL

replacing the unrecognized pattern number position with a ?

on the repo file branch its written to "./data/out.txt"


REST Instrucions

repo rest branch

exec go run main.go

It will run on the port 8080

The endpoint is /ocr and receives raw text data on the body

The response contains a JSON array with the elements presented as such

{"account_number":"796753717","status":"0K"}
{"account_number":"663972126","status":"ERR"}
{"account_number":"?23?06849","status":"ILL"}

the status evaluations are the same as above.



TODO

tests