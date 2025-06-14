package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"WEAKS/testdose/database"
	"WEAKS/testdose/model"
	"WEAKS/testdose/utils"
)

// signup
func Signup(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	Hashpassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(Hashpassword),
		Role:     "user",
	}
	if err := database.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messege": "user registerd"})
}

// signin
func Signin(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		password string `json:"password`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}

	var user model.User
	if err := database.Db.Where("email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// signout

func Signout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	c.JSON(http.StatusAccepted, gin.H{"messege": "logout succesfully"})
}
