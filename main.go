package main

import (
	"bufio"
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
			darwin.Initialise(poolSize)

			for k := range maxEpoch {
				fmt.Println("=========== Epoch", k, "===========")
				fmt.Fprintf(w, "data: Epoch %d\n\n", k)

				darwin.RunEpoch(maxRound)
				darwin.Selection()
				darwin.Log(5)

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

func cli_main() {
	poolSize := 3000
	maxRound := 10

	darwin.Initialise(poolSize)
	for k := range 10 {
		fmt.Println("=========== Epoch", k, "===========")
		darwin.RunEpoch(maxRound)
		darwin.Selection()
		darwin.Log(5)
	}

	// on compte les victoires indépendamment du coût
	// Puis on regroupe les npc par COST pour déterminer qui a la plus de victoire pour un COST donné.
	// Et ensuite on duplique/mute les gagnants de chaque COST. On peut en générer autant qu'il y a de NPC dans un COST donné
	// De cette façon, on ne change pas le profil de puissance globale

	// Il faut visualiser 2 courbes :
	// * le nombre de NPC par COST
	// * le nombre de victoires (total ? moyen ?) par COST

	// L'idée c'est obtenir non pas le meilleur NPC du pool complet mais le meilleur NPC pour un COST donné
	// Donc pour chaque COST donné, on vire (on remplace par des mutants), les NPC qui ont le moins de victoires dans ce COST donné

}
