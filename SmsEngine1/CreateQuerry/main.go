package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Create Query salem XD")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully Connected to the Database")

}
