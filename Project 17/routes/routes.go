package routes

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/controllers"
)

func InitRouters(app *fiber.App) {
	app.Get("/api/v1/blog", controllers.GetBlogs)
	app.Get("/api/v1/blog/:id", controllers.GetBlog)
	app.Post("/api/v1/blog", controllers.ValidateBlogDto, controllers.CreateBlog)
	app.Put("api/v1/blog/:id", controllers.ValidateBlogDto, controllers.UpdateBlog)
	app.Delete("/api/v1/blog/:id", controllers.DeleteBlog)
}