package main

import (
	"fmt"
	"log"
	"os"

	// "log"
	// "os"
	// "github.com/joho/godotenv"
	"github.com/RohitKuwar/go_fiber/routes"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	// app.Use(logger.New())
	routes.Setup(app)

	app.Listen(":" + port)
	fmt.Printf("Server is running on port: %v\n", port)
	log.Printf("Server is running on port: %v\n", port)

}
