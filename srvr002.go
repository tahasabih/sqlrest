package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "statushku1:hkuer123@tcp(85.10.205.173:3306)/statushku")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	fmt.Printf("database open")
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Query("INSERT INTO users (username, password_hash, accountType) VALUES('Tahasa', 'abcd123', 'admin')") // ? = placeholder
	if err != nil {
		panic(err.Error())
		fmt.Printf("could not insert now") // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
}

// Insert square numbers for 0-24 in the database
/*for i := 0; i < 25; i++ {
    _, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}*/
