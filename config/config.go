package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/packsDb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}
