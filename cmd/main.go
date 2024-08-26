package main

import (
	"github.com/gofiber/fiber/v2"
	"project/internal/adapter/config"
	"project/internal/adapter/handler/http"
	"project/internal/adapter/storage/pg"
	"project/internal/adapter/storage/pg/repository"
	"project/internal/core/service"
)

func init() {
	config.Initial()
}

func main() {
	db, err := pg.New(config.GetPostgres())
	if err != nil {
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
