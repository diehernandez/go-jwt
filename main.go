package main

import (
	"fmt"

	"github.com/diehernandez/go-auth/database"
	"github.com/diehernandez/go-auth/routes"
	"github.com/diehernandez/go-auth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// port := viperEnvVariable("PORT")
	urlConnection := utils.ViperEnvVariable("URL_CONNECTION")
	fmt.Printf("viper : %s = %s \n", "URL", urlConnection)

	database.Connect(&urlConnection)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(":8000")
}
