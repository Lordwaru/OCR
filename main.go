package main

import (
	"net/http"

	"github.com/Lordwaru/OCR/docs"
	"github.com/Lordwaru/OCR/internal/ocr"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	router := gin.Default()

	docs.SwaggerInfo.BasePath = ""

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Static("/css", "./web/css")
	router.Static("/fonts", "./web/fonts")
	router.Static("/js", "./web/js")

	router.LoadHTMLGlob("web/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", GetIndex)

	ocr.AddHandlers(router)

	router.Run(":8080")

}

func GetIndex(c *gin.Context) {
	accs := ocr.SelectAllAccounts()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"accounts": accs,
	})
}
