package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
func Update(token string) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/todo")
	query := "UPDATE todo SET can='yes' WHERE token=?"
	res, err := db.Exec(query,token)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowc, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d Güncelledik\n", rowc)
}