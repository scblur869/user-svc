package _data

import (
	"database/sql"
	"fmt"
	"local/user-svc/_models"
)

func GetAllUsers() []_models.EMPLOYEE {

	db, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}

	var employee _models.EMPLOYEE
	var employeeList []_models.EMPLOYEE
	query := "SELECT id, first_name, last_name, email, street_address, gender, department FROM employee"
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var id int
		var first_name string
		var last_name string
		var email string
		var street_address string
		var gender string
		var department string

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &first_name, &last_name, &email, &street_address, &gender, &department)
		if err != nil {
			fmt.Println(err.Error())
		}
		employee.Id = id
		employee.First_Name = first_name
		employee.Last_Name = last_name
		employee.Email = email
		employee.Street_Address = street_address
		employee.Gender = gender
		employee.Department = department

		employeeList = append(employeeList, employee)

	}
	db.Close()
	return employeeList
}

func SelectByParam(column string, param string) []_models.EMPLOYEE {

	db, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}

	var employee _models.EMPLOYEE
	var employeeList []_models.EMPLOYEE
	query := "SELECT id, first_name, last_name, email, street_address, gender, department FROM employee where " + column + " = ?"
	results, err := db.Query(query, param)
	if err != nil {
		fmt.Println(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var id int
		var first_name string
		var last_name string
		var email string
		var street_address string
		var gender string
		var department string

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &first_name, &last_name, &email, &street_address, &gender, &department)
		if err != nil {
			fmt.Println(err.Error())
		}
		employee.Id = id
		employee.First_Name = first_name
		employee.Last_Name = last_name
		employee.Email = email
		employee.Street_Address = street_address
		employee.Gender = gender
		employee.Department = department

		employeeList = append(employeeList, employee)

	}
	db.Close()
	return employeeList
}
