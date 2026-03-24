package logic

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

func PrintBanner() {
	fmt.Println(ColorRed + "  _                              _            __   __               ")
	fmt.Println(" | |    ___  __ _ ___ _ __   __| |___   ___  / _|  \\ \\ / /__  _____ ")
	fmt.Println(" | |   / _ \\/ _` |/ _ \\ '_ \\ / _` / __| / _ \\| |_    \\ V / _ \\/ \\ \\ / /")
	fmt.Println(" | |__|  __/ (_| |  __/ | | | (_| \\__ \\| (_) |  _|    | | (_) \\ V  V / ")
	fmt.Println(" |_____\\___|\\__, |\\___|_| |_|\\__,_|___/ \\___/|_|      |_|\\___/ \\_/\\_/  ")
	fmt.Println("            |___/                                                       " + ColorReset)
	fmt.Println(ColorCyan + "                --- PROJET RED : LE JEU DE RÔLE ---                " + ColorReset)
}

func Introduction() Character {
	ClearScreen()
	fmt.Println(ColorCyan + `
       />_________________________________
[########[]_________________________________>
       \>
` + ColorReset)
	fmt.Println(ColorBlue + "=================================================================" + ColorReset)
	fmt.Println("Il y a de cela des siècles, le monde d'Ynov était un havre de paix.")
	fmt.Println("Gouverné par les anciens Gardiens, la magie et la nature prospéraient.")
	fmt.Println("Mais une ombre s'est abattue sur ces terres. Le Roi Démon s'est éveillé,")
	fmt.Println("et avec lui, des hordes de créatures cauchemardesques ont envahi nos foyers.")
	fmt.Println(ColorBlue + "=================================================================" + ColorReset)
	WaitForInput()

	ClearScreen()
	fmt.Println(ColorPurple + `
           |>>>
           |
       _  _|_  _
      |;|_|;|_|;|
      \\.    .  /
       \\:  .  /
        ||:   |
        ||_   |
` + ColorReset)
	fmt.Println(ColorBlue + "=================================================================" + ColorReset)
	fmt.Println("Aujourd'hui, les châteaux sont en ruines et l'espoir s'amenuise.")
	fmt.Println("Vous vous réveillez au milieu d'un champ de bataille oublié,")
	fmt.Println("la tête lourde, serrant le pommeau de votre arme émoussée.")
	fmt.Println("Une seule pensée résonne dans votre esprit...")
	fmt.Println("Les chroniques racontent l'histoire d'un héros, mais avant")
	fmt.Println("de devenir une légende, tout commence par un nom.")
	fmt.Println(ColorBlue + "=================================================================" + ColorReset)
	fmt.Println()

	player := characterCreation()

	ClearScreen()
	fmt.Println(ColorYellow + `
      ~         ~~          __
          _T      .,,.    ~--~ ^^
    ^^   // \                    ~
         ][O]    ^^      ,-~ ~
      /''-I_I         _II____
   __/_  /   \ ______/ ''   /'\_,__
     | II--'''' \,--:--..,_/,.-{ },
` + ColorReset)
	fmt.Printf("\nBienvenue, %s le %s.\n", player.Name, player.Class)
	fmt.Println("Vous vous relevez péniblement et marchez quelques heures...")
	fmt.Println("Votre aventure commence ici, aux portes du village de Bourg-Palette.")
	WaitForInput()
	return player
}

func Quest1_TheVillage(player *Character) {
	isTrainingDone := false

	for {
		ClearScreen()
		fmt.Println(ColorCyan + `
      .---.
     /_____\
     |  _  |
    _| | | |_
   (___|_|___)  --- CHAPITRE 1 : LE VILLAGE ---
` + ColorReset)
		fmt.Println("Vous vous tenez sur la place centrale du village.")
		fmt.Println("Les villageois vaquent à leurs occupations. Un marchand vous fait signe.")

		if !isTrainingDone {
			fmt.Println("Le maître d'armes vous observe. " + ColorRed + "Vous devriez vous entraîner avant d'aller plus loin." + ColorReset)
		} else {
			fmt.Println("Le forgeron a remarqué vos talents et vous a ouvert les portes de son atelier.")
			fmt.Println("La porte vers la forêt (Zone suivante) est maintenant déverrouillée.")
		}
		fmt.Println()
		fmt.Printf("Joueur : %s (%s) | PV : %s%d/%d%s | Or : %s%d%s\n", player.Name, player.Class, ColorGreen, player.CurrentHP, player.MaxHP, ColorReset, ColorYellow, player.Money, ColorReset)
		printSeparator()
		fmt.Println("1. 📜 Afficher la Fiche Personnage")
		fmt.Println("2. 🎒 Accéder à l'inventaire")
		fmt.Println("3. 💰 Voir le Marchand")
		fmt.Println("4. ⚔️  S'entraîner contre un Gobelin")

		if isTrainingDone {
			fmt.Println("5. ⚒️  Forgeron (Débloqué)")
			fmt.Println("6. 🌲 Quitter le village et entrer dans la Forêt (Chapitre suivant)")
		} else {
			fmt.Println("5. ⚒️  Forgeron (Verrouillé - Entraînez-vous d'abord)")
		}
		fmt.Println("7. ❓ Qui sont-ils ?")
		fmt.Print("\nVotre choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			ClearScreen()
			displayInfo(player)
			WaitForInput()
		case "2":
			accessInventory(player, reader)
		case "3":
			merchantMenu(player, reader)
		case "4":
			fmt.Println("\nVous approchez du maître d'armes...")
			fmt.Println(ColorYellow + "      /| ________________")
			fmt.Println("O|===|* >________________>")
			fmt.Println("      \\|" + ColorReset)
			WaitForInput()
			trainingMenu(player, reader)
			isTrainingDone = true
		case "5":
			if isTrainingDone {
				blacksmithMenu(player, reader)
			} else {
				fmt.Println(ColorRed + "\nLe forgeron grogne : 'Reviens quand tu sauras tenir une arme, bleuaille !'" + ColorReset)
				WaitForInput()
			}
		case "6":
			if isTrainingDone {
				fmt.Println("\nVous rassemblez vos affaires et vous dirigez vers la sortie du village...")
				WaitForInput()
				return
			}
			fallthrough
		case "7":
			fmt.Println("\n🌟 Projet créé par ADAMA DEZE KONATE & Aurélie")
			WaitForInput()
		default:
			fmt.Println("\nChoix invalide, veuillez réessayer.")
			WaitForInput()
		}
	}
}

func Quest2_TheForest(player *Character) {
	ClearScreen()
	fmt.Println(ColorGreen + "\n--- CHAPITRE 2 : LA FORÊT SOMBRE ---" + ColorReset)
	fmt.Println("L'air se rafraîchit à mesure que vous pénétrez sous la canopée épaisse.")
	fmt.Println("Des bruits inquiétants résonnent autour de vous.")
	fmt.Println(ColorGreen + `
           .
          / \
         /   \
        /_____\
       /       \
      /_________\
     /           \
    /_____________\
          |||
          |||
` + ColorReset)
	fmt.Println("\nSoudain, une bête surgit des buissons !")
	WaitForInput()

	StartForestEncounter(player, reader)
}

func Quest3_TheFortress(player *Character) {
	ClearScreen()
	fmt.Println(ColorPurple + "\n--- CHAPITRE 3 : LA FORTERESSE SOMBRE ---" + ColorReset)
	fmt.Println("Après avoir vaincu le loup, vous trouvez un chemin étroit menant à une ancienne forteresse.")
	fmt.Println("Les portes géantes sont entrouvertes et une odeur nauséabonde s'en échappe...")
	fmt.Println(ColorPurple + `
             |>>>
             |
         _  _|_  _
        |;|_|;|_|;|
        \\.    .  /
         \\:  .  /
          ||:   |
          ||_   |
` + ColorReset)
	fmt.Println("\nUn grondement terrible fait trembler le sol. Le gardien des lieux apparaît !")
	WaitForInput()

	StartFortressEncounter(player, reader)
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

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func WaitForInput() {
	fmt.Println(ColorYellow + "\nAppuyez sur Entrée pour continuer..." + ColorReset)
	reader.ReadString('\n')
}

func printSeparator() {
	fmt.Println(ColorCyan + "----------------------------------------" + ColorReset)
}
