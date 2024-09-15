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
		pool = append(pool, buildFighter(4+2*rand.Intn(5), rand.Intn(3), 4+2*rand.Intn(5)))
	}

	for f1 := 0; f1 < poolSize; f1++ {
		for f2 := 0; f2 < f1; f2++ {
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
				fighter1.receiveAttack(fighter2)
				fighter2.receiveAttack(fighter1)
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
		return b.victory - a.victory
	})

	for k := 0; k < 10; k++ {
		fmt.Println(pool[k])
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

	// l'autre solution est de faire des combats random pour faire une analyse de l'effet du coût sur le nombre de victoire, on obtient une courbe
	// Ensuite on compte des points de victoire pondérés en fonction de cette courbe (foireux)

}
