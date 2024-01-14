package main

import (
	"quiz-maker-backend/database"
	"quiz-maker-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("quizmaker.db")
	database.Migrate()

	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	apiRoutes := router.Group("/api/v1")
	routes.QuizRoutes(apiRoutes)
	routes.SubmissionRoutes(apiRoutes)

	return router
}
