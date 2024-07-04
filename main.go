package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"rate-limiting-middleware/config"
)

func main() {
	conf, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	app := fiber.New()

	app.Use("/endpoint1", limiter.New(limiter.Config{
		Max:        conf.Endpoints["/endpoint1"].Limit,
		Expiration: time.Duration(conf.Endpoints["/endpoint1"].IntervalSeconds) * time.Second,
	}))

	app.Use("/endpoint2", limiter.New(limiter. Config{
		Max:        conf.Endpoints["/endpoint2"].Limit,
		Expiration: time.Duration(conf.Endpoints["/endpoint2"].IntervalSeconds) * time.Second,
	}))

	app.Get("/endpoint1", func(c *fiber.Ctx) error {
		return c.SendString("Endpoint 1")
	})

	app.Get("/endpoint2", func(c *fiber.Ctx) error {
		return c.SendString("Endpoint 2")
	})

	log.Fatal(app.Listen(":8080"))
}
