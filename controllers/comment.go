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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	commentRequest := models.CreateCommentRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	comment := models.Comment{
		PhotoId: commentRequest.PhotoId,
		Message: commentRequest.Message,
		UserId:  userID,
	}

	err := db.Debug().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	commentString, _ := json.Marshal(comment)
	commentResponse := models.CreateCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	c.JSON(http.StatusCreated, commentResponse)
}

func GetAllComments(c *gin.Context) {
	db := database.GetDB()

	comments := []models.Comment{}

	err := db.Debug().Preload("User").Order("id asc").Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	commentsString, _ := json.Marshal(comments)
	commentsResponse := []models.GetCommentResponse{}
	json.Unmarshal(commentsString, &commentsResponse)

	c.JSON(http.StatusOK, commentsResponse)
}

func GetOneComment(c *gin.Context) {
	db := database.GetDB()

	comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	err := db.First(&comment, "id = ?", commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	commentString, _ := json.Marshal(comment)
	commentResponse := models.GetCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	c.JSON(http.StatusOK, commentResponse)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	commentRequest := models.UpdateCommentRequest{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserId = userID

	updateString, _ := json.Marshal(commentRequest)
	updateData := models.Comment{}
	json.Unmarshal(updateString, &updateData)

	err := db.Model(&comment).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&comment, comment.ID).Error

	commentString, _ := json.Marshal(comment)
	commentResponse := models.UpdateCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	c.JSON(http.StatusOK, commentResponse)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserId = userID

	err := db.Delete(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
