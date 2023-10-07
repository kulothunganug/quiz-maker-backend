package routes

import (
	"quiz-maker-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SubmissionRoutes(e *gin.RouterGroup) {
	submissionRoutes := e.Group("/submissions")
	submissionRoutes.POST("/", controllers.SubmitAnswer)
	submissionRoutes.GET("/:id", controllers.GetSubmissions)
}
