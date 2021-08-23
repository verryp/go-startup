package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/verryp/go-startup/handler"
	"github.com/verryp/go-startup/repository"
	"github.com/verryp/go-startup/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=root dbname=crowdfund_db port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database connected...")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService()

	fmt.Println(authService.GenerateToken(100))

	userHandler := handler.NewUserHandler(userService, authService)
	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/users/email/check", userHandler.CheckEmailAvailability)
	api.POST("/users/avatar/upload", userHandler.UploadAvatar)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
