package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/kelseyhightower/envconfig"
)

var (
	err error
	app *fiber.App
	cfg Config
)

type Config struct {
	Port           int
	AddressAPI     string `split_words:"true"`
	AddressWebsite string `split_words:"true"`
}

func init() {

	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	app = fiber.New()
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		Handler: errHandler,
	}))
}

func main() {
	setupRoutes()
	serve()
}

func setupRoutes() {
	app.Get("/", rootHandler)

	app.Get("/b", func(c *fiber.Ctx) {
		panic("Something went wrong!")
	})
}

func serve() {
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})
	app.Listen(cfg.Port)
}

func rootHandler(c *fiber.Ctx) {
	a := make(map[string]interface{}, 0)
	a["message"] = "Welcome to Payload Pro"

	l := make(map[string]string, 0)
	l["api"] = cfg.AddressAPI
	l["site"] = cfg.AddressWebsite
	c.JSON(response{
		Data: data{
			Type:       "root",
			Attributes: a,
			Links:      l,
		},
	})
	c.SendStatus(http.StatusOK)
}
