package _models

type EMPLOYEE struct {
	Id             int    `json:"id"`
	First_Name     string `json:"firstname"`
	Last_Name      string `json:"lastname"`
	Street_Address string `json:"streetaddress"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Personel_Id    string `json:"personelid"`
	Department     string `json:"department"`
	Phone          string `json:"phone"`
	Birthdate      string `json:"birthdate"`
	Photo_Id       string `json:"photoid"`
	Attribute_1    string `json:"attribute1"`
	Attribute_2    string `json:"attribute2"`
}

type ListPayload struct {
	Employee []EMPLOYEE `json:"employee"`
	Auth     DeviceAuth `json:"auth"`
}

type DeviceAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
	DeviceIP string `json:"deviceip"`
}

type ID struct {
	Id int `json:"id"`
}
