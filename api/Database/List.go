package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)




type Tag struct {
	Title string `json:"title"`
	No    int    `json:"no"`
	Token string `json:"token"`
	Can   string `json:"can"`
}



func List() ([]Tag, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/todo")
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM todo")
	if err != nil {
		panic(err.Error())
	}
	tags := []Tag{}
	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.Title, &tag.No, &tag.Token,&tag.Can,)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(tag)
		tags = append(tags, tag)
	}

	return tags, nil
}
