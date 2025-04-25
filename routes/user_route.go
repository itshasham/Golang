package routes

import (
	"go-fiber-clean-api/internals/user/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userController *controller.UserController) {
	api := app.Group("/api/users")

	// Correctly use the controller's methods
	api.Get("/", userController.GetUsers)         // Use userController's method
	api.Get("/:id", userController.GetUser)       // Use userController's method
	api.Post("/", userController.CreateUser)      // Use userController's method
	api.Put("/:id", userController.UpdateUser)    // Use userController's method
	api.Delete("/:id", userController.DeleteUser) // Use userController's method
}
