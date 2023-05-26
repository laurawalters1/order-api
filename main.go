package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/laurawalters1/order-api/calculatePacks"
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
	// var totals = calculatePacks.CalculatePacks(int(order.Count))

	context.IndentedJSON(http.StatusOK, "hello")
}

func main() {

	router := gin.Default()
	router.GET("/hi", sayHi)
	router.GET("/place-order", placeOrder)
	router.Run(":3000")
}
