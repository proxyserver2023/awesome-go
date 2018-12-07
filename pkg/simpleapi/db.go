package simpleapi

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root"+"@tcp(172.17.0.2:3306)/test-db")
	CheckErr(err)
	defer db.Close()
	if err != nil {
		fmt.Println("Database Connection succesful")
	}

	return db
}
