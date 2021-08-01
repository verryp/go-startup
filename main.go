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
	dsn := "host=localhost user=crowdfund password=crowdfund dbname=crowdfund_db port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database connected...") 


	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()

}