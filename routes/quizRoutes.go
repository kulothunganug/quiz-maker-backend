package routes

import (
	"quiz-maker-backend/controllers"

	"github.com/gin-gonic/gin"
)

func QuizRoutes(e *gin.RouterGroup) {
	quizRoutes := e.Group("/quizzes")
	quizRoutes.POST("/", controllers.CreateQuiz)
	quizRoutes.GET("/:id", controllers.GetQuiz)
}
