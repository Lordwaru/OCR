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
// @Summary ping example
// @Schemes
// @Description Returns an array with all the accounts, it's status nested by the original encoded base64 string
// @Tags ocr
// @Accept application/json
// @Produce json
// @Success 200 {object} []entity.AccountsJSON "hi"
// @Router /v1/ocr/create-accounts/ [post]
func CreateAccounts(c *gin.Context) {
	PostAccounts(c)
}

// @BasePath /v1/ocr

// OCR Example
// @Summary ping example
// @Schemes
// @Description Returns an array with all the accounts, it's status nested by the original encoded base64 string
// @Tags ocr
// @Accept application/json
// @Produce json
// @Success 200 {object} []entity.AccountsByOriginId "hi"
// @Router /v1/ocr/accounts/ [get]
func GetAccounts(c *gin.Context) {
	GetAllAccounts(c)
}

// @BasePath /v1/ocr

// OCR Example
// @Summary ping example
// @Schemes
// @Description Returns an array with all the accounts, it's status nested by the original encoded base64 string
// @Tags ocr
// @Accept application/json
// @Produce json
// @Success 200 {object} entity.AccountsByOriginId "hi"
// @Router /v1/ocr/accounts/{id} [get]
func GetById(c *gin.Context) {
	SelectAccountsById(c)
}
