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
	query := "SELECT id, first_name, last_name, email, street_address, gender, department, personel_id, phone, birthdate, photo_id, attribute_1, attribute_2 FROM employee"
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
		var personelId string
		var phone string
		var birthdate string
		var photo_id string
		var attribute_1 string
		var attribute_2 string

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &first_name, &last_name, &email, &street_address, &gender, &department, &personelId, &phone, &birthdate, &photo_id, &attribute_1, &attribute_2)
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
		employee.Personel_Id = personelId
		employee.Phone = phone
		employee.Birthdate = birthdate
		employee.Attribute_1 = attribute_1
		employee.Attribute_2 = attribute_2
		employee.Photo_Id = photo_id
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
	query := "SELECT id, first_name, last_name, email, street_address, gender, department, personel_id, phone, birthdate, photo_id, attribute_1, attribute_2  FROM employee where " + column + " = ?"
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
		var personelId string
		var phone string
		var birthdate string
		var photo_id string
		var attribute_1 string
		var attribute_2 string

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &first_name, &last_name, &email, &street_address, &gender, &department, &personelId, &phone, &birthdate, &photo_id, &attribute_1, &attribute_2)
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
		employee.Personel_Id = personelId
		employee.Phone = phone
		employee.Birthdate = birthdate
		employee.Attribute_1 = attribute_1
		employee.Attribute_2 = attribute_2
		employee.Photo_Id = photo_id
		employeeList = append(employeeList, employee)

	}
	db.Close()
	return employeeList
}
