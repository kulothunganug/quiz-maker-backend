package main

import (
	"os"
	"quiz-maker-backend/database"
	"quiz-maker-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("quizmaker.db")
	database.Migrate()

	if os.Getenv("APP_DEV_MODE") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

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
