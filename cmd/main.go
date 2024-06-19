package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"project/internal/adapter/config"
	"project/internal/adapter/handler/http"
	"project/internal/adapter/logs"
	"project/internal/adapter/storage/postgres"
	"project/internal/core/repository"
	"project/internal/core/service"
)

func init() {
	config.Initial()
}

func main() {
	db, err := postgres.New(config.GetPostgresConfig())
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to connect to database: %v", err))
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	server := fiber.New()
	xUser := server.Group("/x-user")
	xUser.Post("/api/create-user", userHandler.CreateUser)

	err = server.Listen(config.GetAppAddr())
	if err != nil {
		panic(err)
		return
	}
}
