package utils

import (
	"BlogAPI/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRouters(app *fiber.App) {
	app.Get("/api/v1/blog/:id", controllers.GetBlog)
	app.Post("/api/v1/blog", controllers.CreateBlog)
	app.Put("api/v1/blog/:id", controllers.UpdateBlog)
	app.Delete("/api/v1/blog/:id", controllers.DeleteBlog)
}
