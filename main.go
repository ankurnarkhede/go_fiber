package main

import (
	"fmt"
	// "log"
	// "os"
	// "github.com/joho/godotenv"
	"github.com/RohitKuwar/go_fiber/routes"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	
	// app.Use(logger.New())
	routes.Setup(app)

	app.Listen(":8000")
	fmt.Println("Server is runnig on port: 8000")
	
}