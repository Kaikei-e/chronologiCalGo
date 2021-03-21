package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var drinkNum = 0


	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")
	router.Static("/assets", "./assets")


	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.tmpl", gin.H{})
	})

	router.POST("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "CaffeCalGo",
			"number": drinkNum, 
		})
	})

	router.Run(":8081")

}
