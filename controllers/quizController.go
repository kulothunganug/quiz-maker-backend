package controllers

import (
	"net/http"
	"quiz-maker-backend/models"
	"quiz-maker-backend/repository"
	"quiz-maker-backend/schemas"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(c *gin.Context) {
	var quiz schemas.Quiz

	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quizId, err := repository.CreateQuiz(quiz)
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
	var quiz models.Quiz
	err := repository.GetQuiz(quizId, true, &quiz)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quiz)
}
