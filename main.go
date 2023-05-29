package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurawalters1/order-api/calculatepacks"
)

type Order struct {
	Count int32 `json:"count"`
}

func placeOrder(context *gin.Context) {
	var order Order
	context.BindJSON(&order)
	var totals = calculatepacks.CalculatePacks(int(order.Count))

	context.IndentedJSON(http.StatusOK, totals)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Control-Allow-Origin")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/place-order", placeOrder)
	router.Run(":3000")
}
