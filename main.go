package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:FT_amir13456679@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err2 := db.Query("INSERT INTO customer(id,name) VALUES ( 2,'test')")

	// if there is an error inserting, handle it
	if err2 != nil {
		panic(err2.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	results, err := db.Query("SELECT id, name FROM customer")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag = newTag()
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}
}
