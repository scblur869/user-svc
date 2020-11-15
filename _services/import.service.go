package _services

import (
	"database/sql"
	"fmt"
	"local/registration-svc/_data"
	"local/registration-svc/_models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const dbpath = "data/db"
const userdb = dbpath + "/user.db"

func ImportRequest(c *gin.Context) {
	var employee _models.EMPLOYEE

	database, err := sql.Open("sqlite3", userdb)
	if err != nil {
		fmt.Println(err)
	}

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	_data.ImportUserRequest()(employee, database)
	c.JSON(http.StatusOK, employee)
}

func RemoveEmployeeById(c *gin.Context) {
	var remove _models.ID
	if err := c.ShouldBindJSON(&remove); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
		database, err := sql.Open("sqlite3", userdb)
		if err != nil {
			fmt.Println(err)
		}
		_data.RemoveEmployeeById(remove.Id, database)
		c.JSON(http.StatusOK, remove)
	}
}

func UpdateEmployeeData(c *gin.Context) {
	var employee _models.EMPLOYEE
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	_data.UpdateEmployeeInfo(employee, employee.Id)
	c.JSON(http.StatusOK, employee)
}

func GetFullEmployeeList(c *gin.Context) {
	employeeList := _data.SelectAll()
	c.JSON(http.StatusOK, employeeList)
}

func GetEmployeeById(c *gin.Context) {
	eId := c.Param("id")
	id, _ := strconv.Atoi(eId)
	employee := _data.SelectByParam("id", id)
	c.JSON(http.StatusOK, employee)
}
