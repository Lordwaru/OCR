package main

import (
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

	ocr.AddHandlers(router)

	router.Run(":8080")

}
