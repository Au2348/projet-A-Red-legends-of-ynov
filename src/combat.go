package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Monster représente un ennemi
type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

// newGoblin crée un Gobelin d'entraînement
func newGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entraînement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

// trainingMenu lance un combat d'entraînement contre un gobelin
func trainingMenu(c *Character, reader *bufio.Reader) {
	fmt.Println("\n⚔️  Un Gobelin d'entraînement surgit !")
	goblin := newGoblin()
	turn := 0

	for {
		turn++
		fmt.Printf("\n--- Tour %d ---\n", turn)
		fmt.Printf("  %s : %d/%d PV\n", c.Name, c.CurrentHP, c.MaxHP)
		fmt.Printf("  %s : %d/%d PV\n", goblin.Name, goblin.CurrentHP, goblin.MaxHP)

		// --- Tour du joueur ---
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
			continue // On ne passe pas au tour du monstre après l'inventaire
		default:
			fmt.Println("  Choix invalide, vous perdez votre tour !")
		}

		// Vérifier si le gobelin est mort
		if goblin.CurrentHP <= 0 {
			fmt.Printf("\n🏆 %s est vaincu ! Vous remportez le combat !\n", goblin.Name)
			return
		}

		// --- Tour du Gobelin ---
		goblinAttaque(c, &goblin, turn)

		// Vérifier si le joueur est mort
		if isDead(c) {
			fmt.Println("Le combat est terminé.")
			return
		}
	}
}

// joueurAttaque gère le choix d'attaque du joueur
func joueurAttaque(c *Character, goblin *Monster, reader *bufio.Reader) {
	fmt.Println("\n  Choisissez votre attaque :")
	fmt.Println("    1. Coup de poing (8 dégâts)")

	// Afficher Boule de Feu seulement si le joueur la connaît
	hasBouleDeFeu := hasSkill(c, "Boule de Feu")
	if hasBouleDeFeu {
		fmt.Println("    2. Boule de Feu  (18 dégâts)")
	}

	fmt.Print("  Votre choix : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		degats := 8
		goblin.CurrentHP -= degats
		fmt.Printf("  👊 %s inflige %d dégâts à %s avec Coup de poing. (%s : %d/%d PV)\n",
			c.Name, degats, goblin.Name, goblin.Name, goblin.CurrentHP, goblin.MaxHP)
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

// goblinAttaque gère l'attaque du gobelin selon son IA
func goblinAttaque(c *Character, goblin *Monster, turn int) {
	var degats int
	var message string

	// Tous les 3 tours, le gobelin frappe 200% plus fort (10 dégâts)
	if turn%3 == 0 {
		degats = 10
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