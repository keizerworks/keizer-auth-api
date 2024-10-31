package server

import (
	"github.com/gofiber/fiber/v2"

	"keizer-auth-api/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "keizer-auth-api",
			AppName:      "keizer-auth-api",
		}),

		db: database.New(),
	}

	return server
}
