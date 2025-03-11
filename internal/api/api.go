package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"pract2/internal/api/middleware"
	"pract2/internal/service"
)

type Routers struct {
	Service *service.Service
}

func NewRouters(r *Routers, token string) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, PUT, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))

	apiGroup := app.Group("/v1", middleware.Authorization(token))
	{
		apiGroup.Post("/create_task", r.Service.Task.CreateTask)
		apiGroup.Get("/get_all_tasks", r.Service.Task.GetAllTasks)
		apiGroup.Get("/get_task/:id", r.Service.Task.GetTaskById)
		apiGroup.Put("/update_task/:id", r.Service.Task.UpdateTaskById)
		apiGroup.Delete("/delete_task/:id", r.Service.Task.DeleteTaskById)
	}

	return app
}
