package main

import (
	"database/sql"
	// "fmt"
	// "io"
	"log"
	// "net/http"
	// "time"

	// "github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

// create a web service
// func main() {

// 	webService := new(restful.WebService)
// 	// create a route  and attach it to the handler in the service
// 	webService.Route(webService.GET("/ping").To(pingTime))

// 	// add the service to the  application
// 	restful.Add(webService)
// 	http.ListenAndServe(":8000", nil)

// }

// func pingTime(req *restful.Request, resp *restful.Response) {

// 	// write to the response

// 	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))

// }

type Book struct{
	Id int
	name string
	author string
}


func main() {
	// Open the database connection
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT, isbn INTEGER, author VARCHAR(64), name VARCHAR(64))")
	if err != nil {
		log.Println("Error in creating the table:", err)
	} else {
		log.Println("Table created: books")
	}
	statement.Exec()

	// Insert data into the table
	statement, err = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error in preparing the insert statement:", err)
	}
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 12345)
	log.Println("Inserted the data into the table")

	// Read data from the table
	rows, err := db.Query("SELECT id, name, author FROM books")
	if err != nil {
		log.Println("Error in querying the data:", err)
	}
	defer rows.Close()

	var tempBook Book
	for rows.Next() {
		err = rows.Scan(&tempBook.Id, &tempBook.name, &tempBook.author)
		if err != nil {
			log.Println("Error in scanning the row:", err)
		}
		log.Printf("ID: %d, Book: %s, Author: %s\n", tempBook.Id, tempBook.name, tempBook.author)
	}

	// Update data in the table
	statement, err = db.Prepare("UPDATE books SET name = ? WHERE id = ?")
	if err != nil {
		log.Println("Error in preparing the update statement:", err)
	}
	statement.Exec("Updated Book: The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in the database")

	// Delete data from the table
	statement, err = db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		log.Println("Error in preparing the delete statement:", err)
	}
	statement.Exec(1)
	log.Println("Successfully deleted the book from the database")
}

