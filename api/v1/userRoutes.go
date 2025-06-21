package v1

import (
	"database/sql"

	"github.com/Aniketyadav44/go-backend-template/internal/handlers"
	"github.com/Aniketyadav44/go-backend-template/internal/services"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.Engine, db *sql.DB) {
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetAllUsers)
	}
}
