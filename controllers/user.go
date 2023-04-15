package controllers

import (
	"encoding/json"
	"my-gram-project/database"
	"my-gram-project/helpers"
	"my-gram-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJson = "application/json"


func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userRequest := models.User{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	userResult := models.User{
		Username: userRequest.Username,
		Email: userRequest.Email,
		Password: helpers.HashPass(userRequest.Password),
		Age: userRequest.Age,
		ProfileImageUrl: "https://linkPhoto.com",
	}

	err := db.Debug().Create(&userResult).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	userString, _ := json.Marshal(userResult)
	userResponse := models.RegisterResponse{}
	json.Unmarshal(userString, &userResponse)

	c.JSON(http.StatusCreated, userResponse)
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userRequest := models.LoginRequest{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	password := userRequest.Password
	user := models.User{}

	err := db.Debug().Where("email = ?", userRequest.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}