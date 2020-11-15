package _models

type EMPLOYEE struct {
	Id             int    `json:"id"`
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
	Street_Address string `json:"street_address"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Personel_Id    string `json:"personel_id"`
	Department     string `json:"department"`
	Phone          string `json:"phone"`
	Birthdate      string `json:"birthdate"`
	Photo_Id       string `json:"photo_id"`
	Attribute_1    string `json:"attribute_1"`
	Attribute_2    string `json:"attribute_2"`
}

type ID struct {
	Id int `json:"id"`
}
