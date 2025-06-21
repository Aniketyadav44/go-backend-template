package main

import (
	"log"

	v1 "github.com/Aniketyadav44/go-backend-template/api/v1"
	"github.com/Aniketyadav44/go-backend-template/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error in loading config: ", err)
	}

	router := gin.Default()
	v1.RegisterRoutes(router, cfg.DB)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("error in starting server: ", err)
	}
}
