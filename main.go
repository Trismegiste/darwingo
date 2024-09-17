package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"main/darwin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/valyala/fasthttp"
)

func main() {
	// Create a new engine
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/css", "./public/css")
	app.Static("/esm", "./public/esm")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		poolSize := c.QueryInt("poolSize")
		maxRound := c.QueryInt("maxRound")
		maxEpoch := c.QueryInt("epoch")

		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("Start simulation")
			world := darwin.BuildWorld(poolSize)

			for k := range maxEpoch {
				fmt.Println("===========", "Epoch", k, "===========")

				world.RunEpoch(maxRound)
				world.Selection()
				world.Log(5)
				stats := world.GetStatPerCost()
				content, _ := json.Marshal(stats)
				fmt.Fprintf(w, "data: %s\n\n", content)

				err := w.Flush()
				if err != nil {
					// Refreshing page in web browser will establish a new
					// SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
			}

			fmt.Fprintf(w, "data: Done\n\n")
		}))

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
