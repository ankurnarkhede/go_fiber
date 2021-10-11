package routes

import (
	// "github.com/RohitKuwar/go_fiber/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Setup(app *fiber.App) {

	app.Get("/knockknock", func(c *fiber.Ctx) error {
		log.Printf("knock knock route")
		return c.SendString("Who is there? 😾")
	})
	// app.Get("/goals", controllers.GetGoals)
	// app.Get("/goals/:id", controllers.GetGoal)
	// app.Post("/goals/", controllers.CreateGoal)
	// app.Patch("/goals/:id", controllers.UpdateGoal)
	// app.Delete("/goals/:id", controllers.DeleteGoal)

}
