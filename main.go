package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHi(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hey!")
}

func main() {
	router := gin.Default()
	router.GET("/hi", sayHi)
	router.Run(":3000")
}
