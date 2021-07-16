package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	//Connection database query
	for i := 0; i < 100; i++ {
		go func(i int) {
			mSql := "select * from user"
			rows, _ := db.Query(mSql)
			rows.Close() // here, if you do not release the connection to the pool, other concurrencies will block after five runs
			fmt.Println(i)
		}(i)

	}

	for {
		time.Sleep(time.Second)
	}
}
