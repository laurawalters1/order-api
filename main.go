package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	context.IndentedJSON(http.StatusOK, order)
}

func main() {
	router := gin.Default()
	router.GET("/hi", sayHi)
	router.GET("/place-order", placeOrder)
	router.Run(":3000")
}
