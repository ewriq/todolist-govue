package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbqery, err := db.Query("SELECT * FROM todo")
	if err != nil {
		panic(err.Error())
	}
	defer dbqery.Close()

	for dbqery.Next() {
		var title string
		var no int
		var token string
		var can string

		err := dbqery.Scan(&title, &no, &token, &can)
		if err != nil {
			panic(err.Error())
		}
        fmt.Println(err)
	}
}
