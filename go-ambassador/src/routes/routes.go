package routes

import (
	"ambassador/src/middlewares"
	"ambassador/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAutenticated := admin.Use(middlewares.IsAuthenticated)
	adminAutenticated.Get("user", controllers.User)
	adminAutenticated.Post("logout", controllers.Logout)


}