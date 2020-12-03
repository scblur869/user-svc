package _services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"local/user-svc/_http"
	"local/user-svc/_models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadUserDataToDevice(c *gin.Context) {

	var employeeList []_models.EMPLOYEE
	var data []_models.DEVICE
	var employee _models.EMPLOYEE
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	var addPersonRequest _models.PersonListRequest
	UserName := "admin"
	PassWord := "admin"
	ts := strconv.Itoa(data[0].TimeStamp)
	strData := []byte(data[0].UUID + ":" + UserName + ":" + PassWord + ":" + ts)
	hasher := md5.New()
	hasher.Write([]byte(strData))
	signedData := hex.EncodeToString(hasher.Sum(nil))

	for _, employee := range employeeList {
		addPersonRequest.Name = "personListRequest"
		addPersonRequest.UUID = data[0].UUID
		addPersonRequest.Session = data[0].UUID + "_" + strconv.Itoa(data[0].TimeStamp)
		addPersonRequest.Timestamp = data[0].TimeStamp
		addPersonRequest.Sign = signedData
		addPersonRequest.Data.Action = "addPerson"
		addPersonRequest.Data.PersonType = 2
		addPersonRequest.Data.FeatureInfo.FeatureType = 0
		addPersonRequest.Data.PersonInfo.PersonCover = 1
		addPersonRequest.Data.PersonInfo.PersonId = strconv.Itoa(employee.Id)
		addPersonRequest.Data.PersonInfo.PersonName = employee.First_Name + " " + employee.Last_Name
		if employee.Gender == "Male" {
			addPersonRequest.Data.PersonInfo.Sex = 1
		} else {
			addPersonRequest.Data.PersonInfo.Sex = 2
		}
		addPersonRequest.Data.PersonInfo.IDCard = employee.Personel_Id
		addPersonRequest.Data.PersonInfo.Nation = "United States"
		addPersonRequest.Data.PersonInfo.Birthday = employee.Birthdate
		addPersonRequest.Data.PersonInfo.Phone = employee.Phone
		addPersonRequest.Data.PersonInfo.Address = employee.Street_Address
		addPersonRequest.Data.PersonInfo.LimitTime = 1
		addPersonRequest.Data.PersonInfo.StartTime = "2019-01-01 00:00:00"
		addPersonRequest.Data.PersonInfo.EndTime = "2021-01-01 00:00:00"
		addPersonRequest.Data.PersonInfo.Label = "HR Verified"
		addPersonRequest.Data.PersonInfo.PersonPhoto = employee.Photo_Id
		addPersonRequest.Data.PersonInfo.FeatureValue = "0"
		_http.AddPersonListToDevice(data[0], addPersonRequest)
	}
}
