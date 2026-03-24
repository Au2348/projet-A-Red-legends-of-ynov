package logic

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// init est appelée automatiquement au lancement du programme pour initialiser l'aléatoire.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Monster regroupe les caractéristiques d'un ennemi (PV, Attaque, etc.).
type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

// newGoblin génère un gobelin d'entraînement dont la puissance s'adapte au niveau du joueur.
func newGoblin(level int) Monster {
	return Monster{
		Name:      fmt.Sprintf("Gobelin d'entraînement (Niv. %d)", level),
		MaxHP:     40 + (level-1)*10,
		CurrentHP: 40 + (level-1)*10,
		Attack:    5 + (level-1)*2,
	}
}

// newWolf génère le monstre redoutable du Chapitre 2.
func newWolf(level int) Monster {
	return Monster{
		Name:      fmt.Sprintf("Loup Enragé (Niv. %d)", level),
		MaxHP:     50 + (level-1)*15,
		CurrentHP: 50 + (level-1)*15,
		Attack:    7 + (level-1)*2,
	}
}

// newTroll génère le boss final très résistant du Chapitre 3.
func newTroll(level int) Monster {
	return Monster{
		Name:      fmt.Sprintf("Troll des Montagnes (Niv. %d)", level),
		MaxHP:     80 + (level-1)*20,
		CurrentHP: 80 + (level-1)*20,
		Attack:    12 + (level-1)*3,
	}
}

// trainingMenu gère le combat tutoriel contre le gobelin dans le village (requis pour avancer).
func trainingMenu(c *Character, reader *bufio.Reader) {
	fmt.Println(ColorRed + `
      ,   ,
      |---|
      |o o|
      |___|
      /   \
     |     |   ⚔️  Un Gobelin d'entraînement surgit !
     |_____|
` + ColorReset)
	goblin := newGoblin(c.Level)
	turn := 0

	for {
		ClearScreen() // Nettoie l'écran à chaque tour
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
			WaitForInput() // Pause avant de quitter le combat
			return
		}

		monsterAttaque(c, &goblin, turn)

		if isDead(c) {
			fmt.Println("Le combat est terminé.")
			WaitForInput()
			return
		}

		// Pause pour laisser le temps au joueur de lire les dégâts du tour
		WaitForInput()
	}
}

// StartForestEncounter gère le combat de l'histoire du Chapitre 2.
func StartForestEncounter(c *Character, reader *bufio.Reader) {
	fmt.Println(ColorRed + `
      ,  ,
      \  /
      (oo)
      /--\   🐺 Un Loup Enragé bondit hors des fourrés !
     /____\
` + ColorReset)
	wolf := newWolf(c.Level)
	turn := 0

	for {
		ClearScreen() // Nettoie l'écran à chaque tour
		turn++
		fmt.Printf(ColorCyan+"\n--- Tour %d ---\n"+ColorReset, turn)
		fmt.Printf("  %s : %s%d/%d PV%s\n", c.Name, ColorGreen, c.CurrentHP, c.MaxHP, ColorReset)
		fmt.Printf("  %s : %s%d/%d PV%s\n", wolf.Name, ColorRed, wolf.CurrentHP, wolf.MaxHP, ColorReset)

		fmt.Println("\n  Que faites-vous ?")
		fmt.Println("    1. Attaquer")
		fmt.Println("    2. Inventaire")
		fmt.Print("  Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			joueurAttaque(c, &wolf, reader)
		case "2":
			accessInventory(c, reader)
			continue
		default:
			fmt.Println("  Choix invalide, vous perdez votre tour !")
		}

		if wolf.CurrentHP <= 0 {
			fmt.Printf(ColorGreen+"\n🏆 %s est vaincu ! Vous remportez le combat !\n"+ColorReset, wolf.Name)
			gainXP(c, 45)
			addInventory(c, "Fourrure de Loup") // Récompense spéciale
			WaitForInput()                      // Pause avant de quitter le combat
			return
		}

		monsterAttaque(c, &wolf, turn)

		if isDead(c) {
			fmt.Println("Le combat est terminé.")
			WaitForInput()
			return
		}

		WaitForInput() // Pause de lecture
	}
}

// StartFortressEncounter gère le combat final et épique du Chapitre 3.
func StartFortressEncounter(c *Character, reader *bufio.Reader) {
	fmt.Println(ColorPurple + `
      .---.
     /     \
    | O   O |
    |   ^   |
     \  -  /    🧌 Un immense Troll des Montagnes bloque le passage !
     /_____\
` + ColorReset)
	troll := newTroll(c.Level)
	turn := 0

	for {
		ClearScreen()
		turn++
		fmt.Printf(ColorCyan+"\n--- Tour %d ---\n"+ColorReset, turn)
		fmt.Printf("  %s : %s%d/%d PV%s\n", c.Name, ColorGreen, c.CurrentHP, c.MaxHP, ColorReset)
		fmt.Printf("  %s : %s%d/%d PV%s\n", troll.Name, ColorPurple, troll.CurrentHP, troll.MaxHP, ColorReset)

		fmt.Println("\n  Que faites-vous ?")
		fmt.Println("    1. Attaquer")
		fmt.Println("    2. Inventaire")
		fmt.Print("  Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			joueurAttaque(c, &troll, reader)
		case "2":
			accessInventory(c, reader)
			continue
		default:
			fmt.Println("  Choix invalide, vous perdez votre tour !")
		}

		if troll.CurrentHP <= 0 {
			fmt.Printf(ColorGreen+"\n🏆 %s est vaincu ! Vous avez libéré la forteresse !\n"+ColorReset, troll.Name)
			gainXP(c, 80)
			addInventory(c, "Peau de Troll")
			WaitForInput()
			return
		}

		monsterAttaque(c, &troll, turn)

		if isDead(c) {
			fmt.Println("Le combat est terminé. La forteresse restera maudite...")
			WaitForInput()
			return
		}

		WaitForInput()
	}
}

// joueurAttaque gère le tour du joueur, affiche ses attaques disponibles et gère la mécanique de coups critiques.
func joueurAttaque(c *Character, m *Monster, reader *bufio.Reader) {
	fmt.Println("\n  Choisissez votre attaque :")
	fmt.Println("    1. Coup de poing (8 dégâts)")

	hasBouleDeFeu := hasSkill(c, "Boule de Feu")
	if hasBouleDeFeu {
		fmt.Println("    2. Boule de Feu  (18 dégâts, 20 Mana)")
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
		m.CurrentHP -= degats
		fmt.Printf("  👊 %s inflige %d dégâts%s à %s avec Coup de poing. (%s : %d/%d PV)\n",
			c.Name, degats, critMsg, m.Name, m.Name, m.CurrentHP, m.MaxHP)
	case "2":
		if hasBouleDeFeu {
			if c.Mana < 20 {
				fmt.Println("  ❌ Pas assez de mana ! L'attaque échoue.")
				return
			}
			c.Mana -= 20
			degats := 18
			m.CurrentHP -= degats
			fmt.Printf(ColorPurple+"  🔥 %s lance une Boule de Feu et inflige %d dégâts à %s ! (%s : %d/%d PV)\n"+ColorReset,
				c.Name, degats, m.Name, m.Name, m.CurrentHP, m.MaxHP)
		} else {
			fmt.Println("  ❌ Vous ne connaissez pas ce sort !")
		}
	default:
		fmt.Println("  Choix invalide, attaque annulée !")
	}
}

// monsterAttaque gère le tour de l'ennemi (celui-ci utilise une attaque chargée tous les 3 tours).
func monsterAttaque(c *Character, m *Monster, turn int) {
	var degats int
	var message string

	if turn%3 == 0 {
		degats = m.Attack * 2
		message = fmt.Sprintf("  💥 %s se déchaîne et inflige %d dégâts à %s ! (Attaque puissante !)",
			m.Name, degats, c.Name)
	} else {
		degats = m.Attack
		message = fmt.Sprintf("  🗡️  %s inflige %d dégâts à %s.",
			m.Name, degats, c.Name)
	}

	c.CurrentHP -= degats
	fmt.Println(message)
	fmt.Printf("  (%s : %d/%d PV)\n", c.Name, c.CurrentHP, c.MaxHP)
}
