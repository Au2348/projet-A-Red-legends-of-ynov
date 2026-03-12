package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Codes couleurs ANSI pour le terminal
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

var reader = bufio.NewReader(os.Stdin)
var player Character

func main() {
	player = characterCreation()

	gameLoop()
}

func gameLoop() {
	for {
		clearScreen()
		fmt.Println(ColorCyan + "========================================" + ColorReset)
		fmt.Println(ColorCyan + "            MENU PRINCIPAL              " + ColorReset)
		fmt.Println(ColorCyan + "========================================" + ColorReset)
		fmt.Printf("Joueur : %s (%s) | PV : %s%d/%d%s | Or : %s%d%s\n", player.Name, player.Class, ColorGreen, player.CurrentHP, player.MaxHP, ColorReset, ColorYellow, player.Money, ColorReset)
		fmt.Println("----------------------------------------")
		fmt.Println("1. Afficher les informations")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Entraînement")
		fmt.Println("6. Quitter")
		printSeparator()
		fmt.Print("Votre choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			clearScreen()
			displayInfo(&player)
			waitForInput()
		case "2":
			accessInventory(&player, reader)
		case "3":
			merchantMenu(&player, reader)
		case "4":
			blacksmithMenu(&player, reader)
		case "5":
			trainingMenu(&player, reader)
		case "6":
			fmt.Println("Fermeture du jeu. À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func characterCreation() Character {
	var name string

	for {
		printSeparator()
		fmt.Print("Entrez le nom de votre héros : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if isValidName(input) {
			name = formatName(input)
			break
		} else {
			fmt.Println("Erreur : Le nom ne doit contenir que des lettres.")
		}
	}

	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println("1. Humain")
		fmt.Println("2. Elfe")
		fmt.Println("3. Nain")
		fmt.Print("Votre choix : ")

		classChoice, _ := reader.ReadString('\n')
		classChoice = strings.TrimSpace(classChoice)

		switch classChoice {
		case "1":
			return newCharacter(name, "Humain")
		case "2":
			return newCharacter(name, "Elfe")
		case "3":
			return newCharacter(name, "Nain")
		default:
			fmt.Println("Classe inconnue, veuillez choisir 1, 2 ou 3.")
		}
	}
}

func isValidName(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatName(s string) string {
	s = strings.ToLower(s)
	return strings.ToUpper(string(s[0])) + s[1:]
}

// clearScreen efface le terminal (Compatible Linux/Mac et Windows récents)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// waitForInput met le jeu en pause jusqu'à ce que le joueur appuie sur Entrée
func waitForInput() {
	fmt.Println(ColorYellow + "\nAppuyez sur Entrée pour continuer..." + ColorReset)
	reader.ReadString('\n')
}

func printSeparator() {
	fmt.Println(ColorCyan + "----------------------------------------" + ColorReset)
}
