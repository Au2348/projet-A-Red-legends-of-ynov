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
		clearScreen() // Nettoie l'écran à chaque tour
		turn++
		fmt.Printf(ColorCyan+"\n--- Tour %d ---\n"+ColorReset, turn)
		fmt.Printf("  %s : %s%d/%d PV%s\n", c.Name, ColorGreen, c.CurrentHP, c.MaxHP, ColorReset)
		fmt.Printf("  %s : %s%d/%d PV%s\n", goblin.Name, ColorRed, goblin.CurrentHP, goblin.MaxHP, ColorReset)

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
			fmt.Printf(ColorGreen+"\n🏆 %s est vaincu ! Vous remportez le combat !\n"+ColorReset, goblin.Name)
			gainXP(c, 30)
			waitForInput() // Pause avant de quitter le combat
			return
		}

		goblinAttaque(c, &goblin, turn)

		if isDead(c) {
			fmt.Println("Le combat est terminé.")
			waitForInput()
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
		critMsg = ColorRed + " (COUP CRITIQUE ! 💥)" + ColorReset
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
			fmt.Printf(ColorPurple+"  🔥 %s lance une Boule de Feu et inflige %d dégâts à %s ! (%s : %d/%d PV)\n"+ColorReset,
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
	time.Sleep(1 * time.Second) // Petite pause pour lire l'attaque de l'ennemi
	fmt.Printf("  (%s : %d/%d PV)\n", c.Name, c.CurrentHP, c.MaxHP)
}