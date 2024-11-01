package main

import (
	"data_app/api"
	"data_app/database"
	"data_app/middleware"
	"data_app/routes"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func init() {
	if err := api.InitConfigs(); err != nil {
		log.Fatalf("An error occured while loading configs: %v", err.Error())
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Can not connect database: %v", err.Error())
	}

	database.InitRedis()
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: api.Config.Api.Prefork,
	})

	// set authentication middleware
	app.Use(middleware.Auth)

	// set cache middleware
	app.Use(cache.New(cache.Config{
		Expiration: time.Second * 10,
	}))

	// set routes
	routes.SetRoutes(app)

	// serve
	app.Listen(fmt.Sprintf("%s:%d", api.Config.Api.Host, api.Config.Api.Port))
}
