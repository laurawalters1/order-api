package config

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func Connection() {
	err := godotenv.Load()

	db, err := sql.Open(os.Getenv("DBCONNTYPE"), os.Getenv("DBCONNNAME"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

}
