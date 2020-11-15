package _data

import (
	"database/sql"
	"fmt"
	"local/user-svc/_models"
	"log"
)

func UpdateEmployeeInfo(employee _models.EMPLOYEE, id int) {

	db, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("Update employee SET first_name =? , last_name =? , street_address =?, gender =?, email =?, personel_id =?, department =?, phone =?, birthdate =?, photo_id =?, attribute_1 =?, attribute_2 =? WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(employee.First_Name, employee.Last_Name, employee.Street_Address, employee.Gender, employee.Email, employee.Personel_Id, employee.Department, employee.Phone, employee.Birthdate, employee.Photo_Id, employee.Attribute_1, employee.Attribute_2, id); err != nil {
		fmt.Println(err)
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
	db.Close()
}
