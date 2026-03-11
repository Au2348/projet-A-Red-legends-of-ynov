package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Character représente la structure de données locale pour le frontend.
// Ta camarade aura probablement une struct similaire côté Backend.
type Character struct {
	Name  string
	Class string
}

// Variables globales pour simplifier la gestion de l'entrée standard
var reader = bufio.NewReader(os.Stdin)
var player Character

func main() {
	// Étape 1 : Création du personnage au lancement
	player = characterCreation()

	// Étape 2 : Boucle principale du jeu
	gameLoop()
}

// gameLoop gère le menu principal
func gameLoop() {
	for {
		printSeparator()
		fmt.Println("--- MENU PRINCIPAL ---")
		fmt.Printf("Joueur : %s (%s)\n", player.Name, player.Class)
		printSeparator()
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
			displayInfo()
		case "2":
			accessInventory()
		case "3":
			merchantMenu()
		case "4":
			blacksmithMenu()
		case "5":
			trainingMenu()
		case "6":
			fmt.Println("Fermeture du jeu. À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

// characterCreation gère la création du nom et le choix de la classe
func characterCreation() Character {
	var p Character

	// 1. Demande et validation du nom
	for {
		printSeparator()
		fmt.Print("Entrez le nom de votre héros : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if isValidName(input) {
			p.Name = formatName(input)
			break
		} else {
			fmt.Println("Erreur : Le nom ne doit contenir que des lettres.")
		}
	}

	// 2. Choix de la classe
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
			p.Class = "Humain"
			return p
		case "2":
			p.Class = "Elfe"
			return p
		case "3":
			p.Class = "Nain"
			return p
		default:
			fmt.Println("Classe inconnue, veuillez choisir 1, 2 ou 3.")
		}
	}
}

// isValidName vérifie si la chaîne ne contient que des lettres
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

// formatName met la première lettre en majuscule et le reste en minuscule
func formatName(s string) string {
	s = strings.ToLower(s)
	return strings.ToUpper(string(s[0])) + s[1:]
}

// --- SOUS-MENUS ---

func merchantMenu() {
	for {
		printSeparator()
		fmt.Println("--- MARCHAND ---")
		fmt.Println("1. Acheter (WIP)")
		fmt.Println("2. Vendre (WIP)")
		fmt.Println("3. Retour au menu principal")
		fmt.Print("Votre choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "3" {
			return // Revient à la boucle gameLoop
		} else {
			fmt.Println("Cette fonctionnalité n'est pas encore disponible.")
		}
	}
}

func blacksmithMenu() {
	for {
		printSeparator()
		fmt.Println("--- FORGERON ---")
		fmt.Println("1. Améliorer équipement (WIP)")
		fmt.Println("2. Retour au menu principal")
		fmt.Print("Votre choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "2" {
			return
		}
		fmt.Println("Le forgeron est en pause.")
	}
}

func trainingMenu() {
	for {
		printSeparator()
		fmt.Println("--- ENTRAÎNEMENT ---")
		fmt.Println("1. Combattre un mannequin (WIP)")
		fmt.Println("2. Retour au menu principal")
		fmt.Print("Votre choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "2" {
			return
		}
		fmt.Println("Zone d'entraînement fermée.")
	}
}

// --- FONCTIONS PLACEHOLDER ---

func displayInfo() {
	printSeparator()
	fmt.Println("--- INFORMATIONS DU PERSONNAGE ---")
	fmt.Printf("Nom : %s\n", player.Name)
	fmt.Printf("Classe : %s\n", player.Class)
	fmt.Println("Niveau : 1 (Backend requis)")
	fmt.Println("PV : 100/100 (Backend requis)")
	fmt.Println("Appuyez sur Entrée pour continuer...")
	reader.ReadString('\n')
}

func accessInventory() {
	printSeparator()
	fmt.Println("--- INVENTAIRE ---")
	fmt.Println("(Vide pour le moment - En attente du Backend)")
	fmt.Println("Appuyez sur Entrée pour continuer...")
	reader.ReadString('\n')
}

// Utilitaires visuels
func printSeparator() {
	fmt.Println(strings.Repeat("-", 30))
}
