package main

import (
	"net/http"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/laurawalters1/order-api/calculatepacks"
	"gorm.io/gorm"
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

// var db *sql.DB

type PackSize struct {
	gorm.Model
	Size int `json:"size"`
}

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/packsDb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	addPackSize := func(context *gin.Context) {
		var packsize PackSize
		context.BindJSON(&packsize)
		db.Create(&packsize)
	}

	// cfg := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "",
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "packsDb",
	// }
	// Get a database handle.

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/place-order", placeOrder)
	router.POST("/add-pack-size", addPackSize)

	router.Run(":3000")
}
