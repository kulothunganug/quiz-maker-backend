package controllers

import (
	"net/http"
	"quiz-maker-backend/models"
	"quiz-maker-backend/repository"
	"quiz-maker-backend/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SubmitAnswer(c *gin.Context) {
	var submissionReq schemas.Submission

	if err := c.ShouldBindJSON(&submissionReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data doesn't match with schema"})
		return
	}
	db := c.MustGet("db")
	dbInstance := db.(gorm.DB)

	submission, err := repository.AddSubmission(&dbInstance, &submissionReq)
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

	db := c.MustGet("db")
	dbInstance := db.(gorm.DB)

	var submissions []models.Submission
	err = repository.GetSubmissions(&dbInstance, uint(quizId), &submissions)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"submissions": submissions})
}
