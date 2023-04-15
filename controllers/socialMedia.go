package controllers

import (
	"encoding/json"
	"my-gram-project/database"
	"my-gram-project/helpers"
	"my-gram-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMediaRequest := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	socialMediaRequest.UserId = userID

	err := db.Debug().Create(&socialMediaRequest).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaString, _ := json.Marshal(socialMediaRequest)
	socialMediaResponse := models.CreateSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusCreated, socialMediaResponse)
}

func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()

	socialMedia := []models.SocialMedia{}

	err := db.Debug().Preload("User").Order("id asc").Find(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := []models.GetSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusOK, socialMediaResponse)
}

func GetOneSocialMedia(c *gin.Context) {
	db := database.GetDB()

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialMedia := models.SocialMedia{}

	err := db.First(&socialMedia, "id = ?", socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialMediaId, errorParam := strconv.Atoi(c.Param("socialMediaId"))
	if errorParam != nil {
		c.AbortWithError(http.StatusBadRequest, errorParam)
		return
	}
	userID := uint(userData["id"].(float64))

	socialMedia := models.SocialMedia{}
	socialMedia.ID = uint(socialMediaId)
	socialMedia.UserId = userID

	err := db.Delete(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMediaRequest := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	socialMedia := models.SocialMedia{}
	socialMedia.ID = uint(socialMediaId)
	socialMedia.UserId = userID

	updateString, _ := json.Marshal(socialMediaRequest)
	updateData := models.SocialMedia{}
	json.Unmarshal(updateString, &updateData)

	err := db.Model(&socialMedia).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&socialMedia, socialMedia.ID).Error

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := models.UpdateSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusOK, socialMediaResponse)
}