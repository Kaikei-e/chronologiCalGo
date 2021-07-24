package main

import (
	"caffecalgo/calculator"
	drinkvalidator "caffecalgo/drinkValidator"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
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
	router.POST("/", drinkvalidator.DrinkNum)

	router.POST("/drinks", drinkvalidator.DrinkNum)
	router.POST("/calculated", func(c *gin.Context) {

		drinkLogs := calculator.DateSorter(calculator.CaffeCalMethod(c))
		calculatedDecaysStructs := calculator.CaffeDecayCals(drinkLogs)
		caledLists := calculatedDecaysStructs.CaffeLists

		// create a new line instance
		kline := charts.NewKLine()

		x := make([]string, 0)
		y := make([]opts.KlineData, 0)
		for i := 0; i < len(caledLists); i++ {
			x = append(x, caledLists[i].DecayTime.Format())
			y = append(y, opts.KlineData{Value: caledLists[i].DecayCaffe})
		}


		kline.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{
				Title: "Caffeine decay in mg and drawn in 5-minute increments",
			}),
			charts.WithXAxisOpts(opts.XAxis{
				SplitNumber: 20,
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Scale: true,
			}),
			charts.WithDataZoomOpts(opts.DataZoom{
				Start:      50,
				End:        100,
				XAxisIndex: []int{0},
			}),
		)

		kline.SetXAxis(x).AddSeries("kline", y)

		

	})



	router.Run(":8085") // listen and serve on 0.0.0.0:8085 (for windows "localhost:8085")
}
