package _data

import (
	"database/sql"
	"fmt"
	"local/registration-svc/_models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ImportUserRequest(employeeData []_models.EMPLOYEE) {

	db, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT INTO employee(first_name, last_name, street_address, gender, email, personel_id, department, phone, birthdate, photo_id, attribute_1, attribute_2 ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for _, employee := range employeeData {
		if _, err := stmt.Exec(employee.First_Name, employee.Last_Name, employee.Street_Address, employee.Gender, employee.Email, employee.Personel_Id, employee.Department, employee.Phone, employee.Birthdate, employee.Photo_Id, employee.Attribute_1, employee.Attribute_2); err != nil {
			fmt.Println(err)
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
	db.Close()
}
