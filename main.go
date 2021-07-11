package main

import (
	"caffecalgo/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently ,"/")
	})
	
	router.GET("/", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "index.html", gin.H{

		})
	})

	router.POST("/", validator.Validator)

	router.POST("/drinks", validator.Validator)


  
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
