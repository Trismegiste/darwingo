package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"main/darwin"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/django/v3"
	"github.com/valyala/fasthttp"
)

func main2() {
	// Create a new engine
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New())

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

			start := time.Now()
			for k := range maxEpoch {
				fmt.Println("===========", "Epoch", k, "===========", time.Since(start))

				world.RunEpoch(maxRound)
				stats := world.GetStatPerCost()
				stats.Epoch = k
				content, _ := json.Marshal(stats)
				fmt.Fprintf(w, "data: %s\n\n", content)
				world.Selection()

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
			world.ExportLdjson("sawo-export.json")
		}))

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}

func rolld6() int {
	roll := 1 + rand.Intn(6)
	switch roll {
	case 4, 5:
		return 1
	case 6:
		return 2
	}
	return 0
}

func rollPoolD6(n int) int {
	sum := 0
	for range n {
		sum += rolld6()
	}
	return sum
}

const maxIter = 10000000

func main() {
	for ndice := 1; ndice <= 10; ndice++ {
		var counter [20]int
		sum := 0
		for range maxIter {
			succ := rollPoolD6(ndice)
			counter[succ]++
			sum += succ
		}
		fmt.Println(ndice, counter)

		var pct [20]float64
		cumul := maxIter
		for idx, count := range counter {
			pct[idx] = float64(100*cumul) / float64(maxIter)
			cumul -= count
		}
		fmt.Println(pct, "\n")
	}
}
