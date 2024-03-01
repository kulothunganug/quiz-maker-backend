package controllers

import (
	"net/http"

	"quiz-maker-backend/models"
	"quiz-maker-backend/repository"
	"quiz-maker-backend/schemas"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateQuiz(c *gin.Context) {
	var quiz schemas.Quiz

	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "Data doesn't match with schema"})
		return
	}

	db := c.MustGet("db")
	dbInstance := db.(*gorm.DB)

	quizId, err := repository.CreateQuiz(dbInstance, quiz)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"message": "Quiz created successfully",
			"quidId":  quizId,
		})
}

func GetQuiz(c *gin.Context) {
	quizId := c.Param("id")
	db := c.MustGet("db")
	dbInstance := db.(*gorm.DB)

	var quiz models.Quiz
	err := repository.GetQuiz(dbInstance, quizId, &quiz)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quiz)
}
