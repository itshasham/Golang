package main

import (
	"go-fiber-clean-api/config"
	"go-fiber-clean-api/internals/user/controller"
	"go-fiber-clean-api/internals/user/repository"
	"go-fiber-clean-api/internals/user/service"
	"go-fiber-clean-api/routes"

	"github.com/gofiber/fiber/v2" // Updated import to the correct Fiber version
)

func main() {
	// Initialize Fiber application
	app := fiber.New()

	// Initialize the database connection
	db := config.InitDB()

	// Initialize the repositories, services, and controllers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo) // Corrected here: `services` instead of `service`
	userController := controller.NewUserController(userService)

	// Register the routes
	routes.UserRoutes(app, userController)

	// Start the Fiber application on port 3000
	app.Listen(":3000")
}
