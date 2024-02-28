// user.go
package routes

import (
	"NeuroNode/controllers"
	"NeuroNode/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	router.Get("/me", middleware.DeserializeUser, controllers.GetMe)
}
