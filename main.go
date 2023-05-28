package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/laurawalters1/order-api/calculatepacks"
)

type Order struct {
	Count int32 `json:"count"`
}

func sayHi(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hey!")
}

func placeOrder(context *gin.Context) {
	var order Order
	context.BindJSON(&order)
	fmt.Println("Count", order.Count)
	var totals = calculatepacks.CalculatePacks(int(order.Count))

	context.IndentedJSON(http.StatusOK, totals)
}

func main() {

	router := gin.Default()
	router.GET("/hi", sayHi)
	router.GET("/place-order", placeOrder)
	router.POST("/place-order", placeOrder)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000.com"},
		AllowMethods: []string{"POST", "GET"},
	}))
	router.Run(":3000")
}
