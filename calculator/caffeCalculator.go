package calculator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CaffeCalMethod(ctx *gin.Context){



	
	ctx.HTML(http.StatusOK, "calculatedPage.html", gin.H{

	})
}