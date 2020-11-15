package _data

import (
	"database/sql"
	"fmt"
	"log"
)

func RemoveEmployeeById(id int) {

	db, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("DELETE FROM employee WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		fmt.Println(err)
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
	db.Close()
}
