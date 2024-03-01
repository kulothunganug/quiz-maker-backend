package main

import (
	"quiz-maker-backend/database"
	"quiz-maker-backend/middlewares"
	"quiz-maker-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	db := database.New("quizmaker.db")
	db.Migrate()

	router := gin.Default()

	apiRoutes := router.Group("/api/v1")
	apiRoutes.Use(middlewares.DatabaseMiddleware(db.Instance))

	routes.QuizRoutes(apiRoutes)
	routes.SubmissionRoutes(apiRoutes)

	return router
}
