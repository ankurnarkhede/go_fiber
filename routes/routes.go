package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/RohitKuwar/go_fiber/controllers"
)

func Setup(app *fiber.App) {

	app.Get("/goals", controllers.GetGoals)
	app.Get("/goals/:id", controllers.GetGoal)
	app.Post("/goals/", controllers.CreateGoal)
	app.Put("/goals/:id", controllers.UpdateGoal)
	app.Delete("/goals/:id", controllers.DeleteGoal)

}
