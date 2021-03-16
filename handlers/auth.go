package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sawasaki-narumi/calcal-api/models"
	"github.com/sawasaki-narumi/calcal-api/utils"
	"net/http"
)

var UserModel = new(models.UserModel)

func userView(user *models.User) gin.H {
	return gin.H{
		"id":   user.ID,
		"name": user.Name,
	}
}

func HandleRegistration(c *gin.Context) {
	// mailとpassword取得
	// 登録が無ければ登録

	auth := &models.Auth{}
	err := c.BindJSON(auth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation Error",
		})
		return
	}

	if auth.Password != auth.ConfirmationPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords doesn't match",
		})
		return
	}

	hash := utils.HashAndSalt([]byte(auth.Password))
	user := &models.User{
		Name:     auth.Name,
		Email:    auth.Email,
		Password: hash,
	}
	err = UserModel.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can`t create the user",
		})
		return
	}

	c.JSON(http.StatusOK, userView(user))
}

func HandleLogin(c *gin.Context) {
	auth := &models.Auth{}

	err := c.BindJSON(auth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid post data",
		})
	}

	user := UserModel.FindByEmail(auth.Email)
	if user.ID == 0 || utils.CompareHashAndPassword(user.Password, auth.Password) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, userView(user))
}
