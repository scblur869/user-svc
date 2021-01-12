package _services

import (
	"fmt"
	"local/user-svc/_http"
	"local/user-svc/_models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonListOnDevice(c *gin.Context) {
	auth := _models.DeviceAuth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	employeeList := _http.GetPersonListFromDevice(auth)
	c.JSON(http.StatusOK, employeeList)
}
