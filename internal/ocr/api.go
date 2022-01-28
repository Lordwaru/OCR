package ocr

import (
	"github.com/gin-gonic/gin"
)

func AddHandlers(router *gin.Engine) {

	v1 := router.Group("/v1/ocr")
	{
		v1.POST("/create-accounts", CreateAccounts)
		v1.GET("/accounts/", GetAccounts)
		v1.GET("/accounts/:id", GetById)
	}

}

// @BasePath /v1/ocr

// OCR Example
// @Summary Posts a OCR number list
// @Schemes
// @Description Creates a resource for the OCR number list and returns an array with all the accounts, it's status nested by the original encoded base64 string
// @Tags ocr
// @Accept application/json
// @Produce json
// @Success 200 {object} []entity.AccountsJSON "success"
// @Router /v1/ocr/create-accounts/ [post]
func CreateAccounts(c *gin.Context) {
	PostAccounts(c)
}

// @BasePath /v1/ocr

// OCR Example
// @Summary Returns all the accounts
// @Schemes
// @Description Returns an array with all the accounts and it's status, nested by the original encoded base64 string
// @Tags ocr
// @Accept application/json
// @Produce json
// @Success 200 {object} []entity.AccountsByOriginId "success"
// @Router /v1/ocr/accounts/ [get]
func GetAccounts(c *gin.Context) {
	GetAllAccounts(c)
}

// @BasePath /v1/ocr

// OCR Example
// @Summary Returns accounts by the resource id
// @Schemes
// @Description Returns a json object with the id, ocr data encoded in base64, and an array for the accounts and its status
// @Tags ocr
// @Accept application/json
// @Produce json
// @Param        id   path      int  true  "id"
// @Success 200 {object} entity.AccountsByOriginId "success"
// @Router /v1/ocr/accounts/{id} [get]
func GetById(c *gin.Context) {
	SelectAccountsById(c)
}
