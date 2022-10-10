package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

const PORT = 3000

func main() {
	engine := handlebars.New("./web/views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Go + Frontend + Live reload example",
		})
	})

	app.Static("/", "web/static")

	go func() {
		err := app.Listen(fmt.Sprintf(":%d", PORT))
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received

	if app == nil {
		log.Fatal("Server was not started")
	}

	log.Println("Stopping server...")
	_ = app.Shutdown()
}
