package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func newGoblin(level int) Monster {
	return Monster{
		Name:      fmt.Sprintf("Gobelin d'entraînement (Niv. %d)", level),
		MaxHP:     40 + (level-1)*10,
		CurrentHP: 40 + (level-1)*10,
		Attack:    5 + (level-1)*2,
	}
}

func trainingMenu(c *Character, reader *bufio.Reader) {
	fmt.Println("\n⚔️  Un Gobelin d'entraînement surgit !")
	goblin := newGoblin(c.Level)
	rand.Seed(time.Now().UnixNano())
	turn := 0

	for {
		turn++
		fmt.Printf("\n--- Tour %d ---\n", turn)
		fmt.Printf("  %s : %d/%d PV\n", c.Name, c.CurrentHP, c.MaxHP)
		fmt.Printf("  %s : %d/%d PV\n", goblin.Name, goblin.CurrentHP, goblin.MaxHP)

		fmt.Println("\n  Que faites-vous ?")
		fmt.Println("    1. Attaquer")
		fmt.Println("    2. Inventaire")
		fmt.Print("  Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			joueurAttaque(c, &goblin, reader)
		case "2":
			accessInventory(c, reader)
			continue
		default:
			fmt.Println("  Choix invalide, vous perdez votre tour !")
		}

		if goblin.CurrentHP <= 0 {
			fmt.Printf("\n🏆 %s est vaincu ! Vous remportez le combat !\n", goblin.Name)
			gainXP(c, 30)
			return
		}

		goblinAttaque(c, &goblin, turn)

		if isDead(c) {
			fmt.Println("Le combat est terminé.")
			return
		}
	}
}

func joueurAttaque(c *Character, goblin *Monster, reader *bufio.Reader) {
	fmt.Println("\n  Choisissez votre attaque :")
	fmt.Println("    1. Coup de poing (8 dégâts)")

	hasBouleDeFeu := hasSkill(c, "Boule de Feu")
	if hasBouleDeFeu {
		fmt.Println("    2. Boule de Feu  (18 dégâts)")
	}

	fmt.Print("  Votre choix : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	critRate := 20
	isCrit := rand.Intn(100) < critRate
	critMultiplier := 1
	critMsg := ""
	if isCrit {
		critMultiplier = 2
		critMsg = " (COUP CRITIQUE ! 💥)"
	}

	switch input {
	case "1":
		degats := 8 * critMultiplier
		goblin.CurrentHP -= degats
		fmt.Printf("  👊 %s inflige %d dégâts%s à %s avec Coup de poing. (%s : %d/%d PV)\n",
			c.Name, degats, critMsg, goblin.Name, goblin.Name, goblin.CurrentHP, goblin.MaxHP)
	case "2":
		if hasBouleDeFeu {
			degats := 18
			goblin.CurrentHP -= degats
			fmt.Printf("  🔥 %s lance une Boule de Feu et inflige %d dégâts à %s ! (%s : %d/%d PV)\n",
				c.Name, degats, goblin.Name, goblin.Name, goblin.CurrentHP, goblin.MaxHP)
		} else {
			fmt.Println("  ❌ Vous ne connaissez pas ce sort !")
		}
	default:
		fmt.Println("  Choix invalide, attaque annulée !")
	}
}

func goblinAttaque(c *Character, goblin *Monster, turn int) {
	var degats int
	var message string

	if turn%3 == 0 {
		degats = goblin.Attack * 2
		message = fmt.Sprintf("  💥 %s se déchaîne et inflige %d dégâts à %s ! (Attaque puissante !)",
			goblin.Name, degats, c.Name)
	} else {
		degats = goblin.Attack
		message = fmt.Sprintf("  🗡️  %s inflige %d dégâts à %s.",
			goblin.Name, degats, c.Name)
	}

	c.CurrentHP -= degats
	fmt.Println(message)
	fmt.Printf("  (%s : %d/%d PV)\n", c.Name, c.CurrentHP, c.MaxHP)
}