package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Insert(token, title string) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := "INSERT INTO todo (title, can, token) VALUES (?, ?, ?)"
	res, err := db.Exec(query, title, "no", token)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowc, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d Added !\n", rowc)
}
