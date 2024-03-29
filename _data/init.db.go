package _data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const dbpath = "data/db"
const userdb = dbpath + "/users.db"

func InitializeUserDatabase() {
	_, err := os.Stat(dbpath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dbpath, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	database, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}
	query := `CREATE TABLE IF NOT EXISTS employee(
		   id INTEGER PRIMARY KEY AUTOINCREMENT,
			 first_name text,
			 last_name text,
			 street_address text,
			 gender text,
			 email text,
			 personel_id text,
			 department text,
			 phone text,
			 birthdate text,
			 photo_id text,
			 attribute_1 text,
			 attribute_2 text,
			 created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			 updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

	stmt, err :=
		database.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	stmt.Exec()
	database.Close()
}
