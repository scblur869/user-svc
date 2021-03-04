package main

// capturetime (datetime), temperature (string), mask (boolean), picture (base 64 string)
import (
	"local/user-svc/_data"
	"local/user-svc/_services"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	allowedHost := os.Getenv("ALLOWED")
	_data.InitializeUserDatabase()
	// gin.SetMode(gin.ReleaseMode)
	// routes
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedHost, "http://localhost:4200"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Accept", "X-Requested-With", "Content-Type", "Authorization", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/api/v1/get-all", _services.GetFullEmployeeList)
	r.GET("/api/v1/get-by-id/:id", _services.GetEmployeeById)
	r.POST("/api/v1/remove-user", _services.RemoveEmployeeById) // accepts registration request from the camera
	r.POST("/api/v1/add-user", _services.ImportRequest)
	r.POST("/api/v1/update-user", _services.UpdateEmployeeData)
	r.POST("/api/v1/get-personlist", _services.GetPersonListOnDevice)  // gets the personlist from the device
	r.POST("/api/v1/add-personlist", _services.UploadUserDataToDevice) // add person list to device
	r.Run("0.0.0.0:8082")                                              // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
