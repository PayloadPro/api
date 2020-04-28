package main

import (
	"log"

	"github.com/PayloadPro/api/handlers"
	"github.com/PayloadPro/api/responses"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/kelseyhightower/envconfig"
)

var (
	err  error
	app  *fiber.App
	conf Config
)

type Config struct {
	Port int
}

func init() {

	err = envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process("", &responses.Conf)
	if err != nil {
		log.Fatal(err)
	}

	app = fiber.New()
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		Handler: responses.ErrHandler,
	}))
}

func main() {
	setupRoutes()
	serve()
}

func setupRoutes() {
	app.Get("/", handlers.Root)
	app.Post("/bins", handlers.CreateBin)
	app.Get("/bins", handlers.GetBins)
	app.Get("/bins/:id", handlers.GetBin)

	app.Get("/b", func(c *fiber.Ctx) {
		panic("Something went wrong!")
	})
}

func serve() {
	app.Use(responses.NotFound)
	app.Listen(conf.Port)
}
