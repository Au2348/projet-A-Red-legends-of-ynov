package logic

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func addInventory(c *Character, item string) bool {
	if len(c.Inventory) >= c.InventoryMax {
		fmt.Printf("❌ Inventaire plein ! (%d/%d) Impossible d'ajouter : %s\n", len(c.Inventory), c.InventoryMax, item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("✅ %s ajouté à l'inventaire. (%d/%d)\n", item, len(c.Inventory), c.InventoryMax)
	return true
}

func removeFromInventory(c *Character, item string) bool {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func countItem(c *Character, item string) int {
	count := 0
	for _, v := range c.Inventory {
		if v == item {
			count++
		}
	}
	return count
}

func accessInventory(c *Character, reader *bufio.Reader) {
	ClearScreen()
	for {
		fmt.Println(ColorBlue + `
      _______
     /       \
    | []   [] |  ========== INVENTAIRE ==========
    |         |
     \_______/
` + ColorReset)
		if len(c.Inventory) == 0 {
			fmt.Println("  (inventaire vide)")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("  %d. %s\n", i+1, item)
			}
		}
		fmt.Printf("  Places : %d / %d\n", len(c.Inventory), c.InventoryMax)
		fmt.Println("================================")
		fmt.Println("  1. Utiliser une Potion de Vie")
		fmt.Println("  2. Retour")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			if removeFromInventory(c, "Potion de Vie") {
				heal := 50
				c.CurrentHP += heal
				if c.CurrentHP > c.MaxHP {
					c.CurrentHP = c.MaxHP
				}
				fmt.Printf(ColorGreen+"🧪 Vous buvez une Potion de Vie et récupérez %d PV. (PV : %d/%d)\n"+ColorReset, heal, c.CurrentHP, c.MaxHP)
			} else {
				fmt.Println("❌ Vous n'avez pas de Potion de Vie.")
			}
		case "2":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func merchantMenu(c *Character, reader *bufio.Reader) {
	ClearScreen()
	for {
		fmt.Println(ColorYellow + `
       _____
      /  $  \
     |  $$$  |  ========== MARCHAND ==========
      \_____/
` + ColorReset)
		fmt.Printf("  Votre or : %s%d pièces%s\n", ColorYellow, c.Money, ColorReset)
		fmt.Println("------------------------------")
		fmt.Println("  Potions :")
		fmt.Println("    1. Potion de Vie        (3 po)  - Rend 50 PV")
		fmt.Println("    2. Potion de Poison      (6 po)  - Inflige 10 dégâts/sec pendant 3s")
		fmt.Println("  Sorts :")
		fmt.Println("    3. Livre : Boule de Feu (25 po) - Apprend le sort Boule de Feu")
		fmt.Println("  Matériaux :")
		fmt.Println("    4. Fourrure de Loup      (4 po)")
		fmt.Println("    5. Peau de Troll         (7 po)")
		fmt.Println("    6. Cuir de Sanglier      (3 po)")
		fmt.Println("    7. Plume de Corbeau      (1 po)")
		fmt.Println("  Améliorations :")
		fmt.Println("    8. Sac à dos (+10 places) (30 po)")
		fmt.Println("  ---------------------")
		fmt.Println("    9. Quitter le marchand")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			achat(c, "Potion de Vie", 3)
		case "2":
			achatPoison(c)
		case "3":
			if c.Money >= 25 {
				if hasSkill(c, "Boule de Feu") {
					fmt.Println("❌ Vous connaissez déjà ce sort.")
				} else {
					c.Money -= 25
					c.Skills = append(c.Skills, "Boule de Feu")
					fmt.Println("📖 Vous apprenez le sort Boule de Feu !")
				}
			} else {
				fmt.Println("❌ Pas assez d'or.")
			}
		case "4":
			achat(c, "Fourrure de Loup", 4)
		case "5":
			achat(c, "Peau de Troll", 7)
		case "6":
			achat(c, "Cuir de Sanglier", 3)
		case "7":
			achat(c, "Plume de Corbeau", 1)
		case "8":
			UpgradeInventorySlot(c)
		case "9":
			fmt.Println("À bientôt, aventurier !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func achat(c *Character, item string, prix int) {
	if c.Money < prix {
		fmt.Printf("❌ Pas assez d'or. (Vous avez %d po, coût : %d po)\n", c.Money, prix)
		return
	}
	if addInventory(c, item) {
		c.Money -= prix
		fmt.Printf(ColorGreen+"💰 Vous achetez %s pour %d po. Or restant : %d po.\n"+ColorReset, item, prix, c.Money)
	}
}

func achatPoison(c *Character) {
	if c.Money < 6 {
		fmt.Printf("❌ Pas assez d'or. (Vous avez %d po, coût : 6 po)\n", c.Money)
		return
	}
	c.Money -= 6
	fmt.Println("💰 Vous achetez une Potion de Poison pour 6 po.")
	fmt.Println("☠️  Vous l'avalez par erreur ! Le poison commence à agir...")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("  💀 Le poison inflige 10 dégâts ! (PV : %d/%d)\n", c.CurrentHP, c.MaxHP)
	}
	fmt.Println("Le poison s'est dissipé.")
	isDead(c)
}

func UpgradeInventorySlot(c *Character) {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("❌ Vous avez déjà atteint le maximum d'améliorations d'inventaire.")
		return
	}
	if c.Money < 30 {
		fmt.Printf("❌ Pas assez d'or. (Vous avez %d po, coût : 30 po)\n", c.Money)
		return
	}
	c.Money -= 30
	c.InventoryMax += 10
	c.InventoryUpgrades++
	fmt.Printf(ColorGreen+"🎒 Amélioration achetée ! Votre inventaire a maintenant %d places maximum. (%d/3 améliorations)\n"+ColorReset, c.InventoryMax, c.InventoryUpgrades)
}

func blacksmithMenu(c *Character, reader *bufio.Reader) {
	ClearScreen()
	for {
		fmt.Println(ColorRed + `
      _______
     [_______]
      |     |    ========== FORGE ==========
      |_____|
` + ColorReset)
		fmt.Printf("  Votre or : %s%d pièces%s (coût forge : +5 po par objet)\n", ColorYellow, c.Money, ColorReset)
		fmt.Println("---------------------------")
		fmt.Println("  1. Chapeau de l'aventurier  (5 po + 1 Plume + 1 Cuir)     → +10 PV Max, slot Tête")
		fmt.Println("  2. Tunique de l'aventurier  (5 po + 2 Fourrures + 1 Peau)  → +25 PV Max, slot Torse")
		fmt.Println("  3. Bottes de l'aventurier   (5 po + 1 Fourrure + 1 Cuir)   → +15 PV Max, slot Pieds")
		fmt.Println("  4. Retour")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			forger(c, "Chapeau de l'aventurier", 5, 10,
				[]string{"Plume de Corbeau", "Cuir de Sanglier"},
				"Tete")
		case "2":
			forger(c, "Tunique de l'aventurier", 5, 25,
				[]string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"},
				"Torse")
		case "3":
			forger(c, "Bottes de l'aventurier", 5, 15,
				[]string{"Fourrure de Loup", "Cuir de Sanglier"},
				"Pieds")
		case "4":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func forger(c *Character, itemName string, cout int, bonusHP int, materiaux []string, slot string) {
	if c.Money < cout {
		fmt.Printf("❌ Pas assez d'or. (Vous avez %d po, coût : %d po)\n", c.Money, cout)
		return
	}

	besoins := make(map[string]int)
	for _, m := range materiaux {
		besoins[m]++
	}
	for mat, quantite := range besoins {
		if countItem(c, mat) < quantite {
			fmt.Printf("❌ Il vous manque des matériaux : %dx %s.\n", quantite, mat)
			return
		}
	}

	for mat, quantite := range besoins {
		for i := 0; i < quantite; i++ {
			removeFromInventory(c, mat)
		}
	}
	c.Money -= cout

	var ancien string
	switch slot {
	case "Tete":
		ancien = c.Equipment.Tete
		c.Equipment.Tete = itemName
	case "Torse":
		ancien = c.Equipment.Torse
		c.Equipment.Torse = itemName
	case "Pieds":
		ancien = c.Equipment.Pieds
		c.Equipment.Pieds = itemName
	}

	if ancien != "" {
		fmt.Printf("🔄 %s retiré du slot et remis dans l'inventaire.\n", ancien)
		addInventory(c, ancien)
	}

	c.MaxHP += bonusHP
	c.CurrentHP += bonusHP
	fmt.Printf("⚒️  %s fabriqué et équipé ! +%d PV Max.\n", itemName, bonusHP)
}
