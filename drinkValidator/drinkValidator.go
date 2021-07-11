package drinkvalidator

import (
	//"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CaffeLogger struct{
	Number int `json:"numOfDrinks"`
	Method int `json:"calMethods"`
	CaffeineMg int `json:"caffeMg"`
	Amount int `json:"amount"`
	Datetime string
}

type CaffeLogs struct{
	CaffeList []CaffeLogger
}



func DrinkNum(ctx *gin.Context){
	numOfDrinksStr := ctx.PostForm("numOfDrinks")

	numOfDrinks, err := strconv.Atoi(numOfDrinksStr)
	if err != nil {
		log.Fatal(err.Error())

		invalidValue := true
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"invalidValue": invalidValue,
		})
	}

	if numOfDrinks > 10 {
		invalidValue := true
		
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"invalidValue": invalidValue,
		})
	}

	logList := []CaffeLogger{}

	for i := 0; i < numOfDrinks; i++ {
		logList = append(logList, CaffeLogger{i, 1, 0, 0, time.Now().Format("2006/01/23 23:45")})
	}

	

	ctx.HTML(http.StatusOK, "drinks.html", gin.H{
		"numOfDrinks": numOfDrinks,
		"logList": logList,
	})
}



