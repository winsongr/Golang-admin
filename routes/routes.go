package routes

import (
	"goadmin/controllers"
	"goadmin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)

	app.Get("/users", controllers.AllUsers)
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}
