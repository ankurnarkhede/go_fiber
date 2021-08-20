package routes

import (
	"github.com/RohitKuwar/go_fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/goals", controllers.GetGoals)
	app.Get("/goals/:id", controllers.GetGoal)
	app.Post("/goals/", controllers.CreateGoal)
	app.Patch("/goals/:id", controllers.UpdateGoal)
	app.Delete("/goals/:id", controllers.DeleteGoal)

}
