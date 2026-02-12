package main

import (
	"context"
	"log"

	obs "github.com/64366126/Lib_Test"
	obsfiber "github.com/64366126/Lib_Test/middleware/fiber"
	"github.com/gofiber/fiber/v3"
)

func main() {
	shutdown, err := obs.InitFromEnv(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown(context.Background())

	app := fiber.New()
	app.Use(obsfiber.Middleware())

	app.Get("/ping", func(c fiber.Ctx) error {
		ctx := c.Context()
		obs.L(ctx).Info("ping")
		_, span := obs.Tracer().Start(ctx, "ping.handler")
		defer span.End()
		return c.SendString("pong")
	})

	app.Get("/context", func(c fiber.Ctx) error {
		ctx := c.Context()
		corr := obs.CorrelationFromContext(ctx)
		obs.L(ctx).Info("context snapshot")
		return c.JSON(corr)
	})

	log.Fatal(app.Listen(":8080"))
}
