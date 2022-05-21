package main

import (
	"BlogAPI/config"
	"BlogAPI/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	config.ManageMigrations()
	app := fiber.New()
	utils.InitRouters(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalln("Router HTTP error: ", err)
	}
}
