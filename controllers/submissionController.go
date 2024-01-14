package controllers

import (
	"net/http"
	"quiz-maker-backend/models"
	"quiz-maker-backend/repository"
	"quiz-maker-backend/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SubmitAnswer(c *gin.Context) {
	var submissionReq schemas.Submission

	if err := c.ShouldBindJSON(&submissionReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submission, err := repository.AddSubmission(&submissionReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{"message": "Answer submitted successfully", "score": submission.Score})
}

func GetSubmissions(c *gin.Context) {
	quizIdStr := c.Param("id")
	quizId, err := strconv.ParseUint(quizIdStr, 10, 32)
	var submissions []models.Submission
	err = repository.GetSubmissions(uint(quizId), &submissions)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"submissions": submissions})
}
