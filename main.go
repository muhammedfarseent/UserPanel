package main

import (
	"github.com/gin-gonic/gin"

	"WEAKS/testdose/controller"
	"WEAKS/testdose/database"
)

func main() {
	database.ConnectDB()
	r := gin.Default()

	r.POST("Signup", controller.Signup)
	r.POST("Sign", controller.Signin)
	r.POST("Signout", controller.Signout)
	r.Run(":8080")

}
