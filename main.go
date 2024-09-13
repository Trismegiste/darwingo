package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func defunc_main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}

func main() {
	poolSize := 1000
	maxRound := 10
	var pool []Fighter

	for k := 0; k < poolSize; k++ {
		pool = append(pool, Fighter{fighting: 4 + 2*rand.Intn(5), parryBonus: rand.Intn(3)})
	}

	for f1 := 0; f1 < poolSize; f1++ {
		for f2 := 0; f2 < poolSize; f2++ {
			if f1 == f2 {
				continue
			}

			// initialise fight
			//	fmt.Println("Fighting", f1, "versus", f2)
			fighter1 := &pool[f1]
			fighter2 := &pool[f2]
			fighter1.reset()
			fighter2.reset()
			round := 0

			// fight
			//	fmt.Println("Fighting", fighter1, "versus", fighter2)
			for !fighter1.isDead() && !fighter2.isDead() && round < maxRound {
				fighter1.attack(fighter2)
				fighter2.attack(fighter1)
				round++
			}

			// aftermath
			if !fighter1.isDead() && fighter2.isDead() {
				//		fmt.Println("Fighter1 has won")
				fighter1.incVictory()
			}
			if !fighter2.isDead() && fighter1.isDead() {
				//	fmt.Println("Fighter2 has won")
				fighter2.incVictory()
			}
		}
	}

	// stat on current epoch
	slices.SortFunc(pool, func(a, b Fighter) int {
		return b.winning - a.winning
	})

	for k := 0; k < 10; k++ {
		fmt.Println(pool[k])
	}
}
