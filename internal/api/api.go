package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"pract2/internal/api/middleware"
	"pract2/internal/service"
)

type Routers struct {
	Service service.IService
}

// создание маршрутов для нашего роутера

func NewRouters(r *Routers, token string) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, PUT, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))

	auth := app.Group("/auth")
	auth.Post("/sing-up", r.Service.GetUserService().SingUp)
	auth.Get("/sing-in", r.Service.GetUserService().SingIn)

	apiGroup := app.Group("/v1", middleware.Authorization(token)) // middleware
	{
		apiGroup.Post("/create_task", r.Service.GetTaskService().CreateTask)           // создание таски
		apiGroup.Get("/get_all_tasks", r.Service.GetTaskService().GetAllTasks)         // взять все таски
		apiGroup.Get("/get_task/:id", r.Service.GetTaskService().GetTaskById)          // взять таску по id
		apiGroup.Put("/update_task/:id", r.Service.GetTaskService().UpdateTaskById)    // обновить таску
		apiGroup.Delete("/delete_task/:id", r.Service.GetTaskService().DeleteTaskById) // удалить таску
		apiGroup.Delete("/deleteUser", r.Service.GetUserService().DeleteUser)          // пользователь удаляет сам себя
	}

	return app
}
