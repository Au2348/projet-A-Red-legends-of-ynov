package logic

import "fmt"

// Equipment contient les objets actuellement équipés par le joueur.
type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

// Character regroupe toutes les statistiques, l'inventaire et les données du joueur.
type Character struct {
	Name              string
	Class             string
	MaxHP             int
	CurrentHP         int
	Level             int
	CurrentXP         int
	MaxXP             int
	Inventory         []string
	InventoryMax      int
	InventoryUpgrades int
	Mana              int
	MaxMana           int
	Money             int
	Skills            []string
	Equipment         Equipment
}

// newCharacter initialise un héros avec des statistiques de base spécifiques selon sa classe.
func newCharacter(name string, class string) Character {
	c := Character{
		Name:         name,
		Class:        class,
		Level:        1,
		CurrentXP:    0,
		MaxXP:        100,
		Money:        100,
		InventoryMax: 10,
		Skills:       []string{"Coup de poing"},
		Inventory:    []string{},
		Equipment:    Equipment{},
	}

	switch class {
	case "Humain":
		c.MaxHP = 100
		c.CurrentHP = 50
		c.MaxMana = 30
	case "Elfe":
		c.MaxHP = 80
		c.CurrentHP = 40
		c.MaxMana = 50
	case "Nain":
		c.MaxHP = 120
		c.CurrentHP = 60
		c.MaxMana = 20
	default:
		c.MaxHP = 100
		c.CurrentHP = 50
		c.MaxMana = 30
	}
	c.Mana = c.MaxMana

	return c
}

// displayInfo affiche de manière formatée (avec de l'ASCII Art) la fiche complète du personnage.
func displayInfo(c *Character) {
	fmt.Println(ColorCyan + `
       /\
      /  \    ========== FICHE PERSONNAGE ==========
     |    |
     |    |
    /| |\ |
    \| |/ |
     |    |
     |____|
` + ColorReset)
	fmt.Printf("  Nom    : %s%s%s\n", ColorPurple, c.Name, ColorReset)
	fmt.Printf("  Classe : %s\n", c.Class)
	fmt.Printf("  Niveau : %s%d%s (XP : %d / %d)\n", ColorCyan, c.Level, ColorReset, c.CurrentXP, c.MaxXP)
	fmt.Printf("  PV     : %s%d / %d%s\n", ColorGreen, c.CurrentHP, c.MaxHP, ColorReset)
	fmt.Printf("  Mana   : %s%d / %d%s\n", ColorBlue, c.Mana, c.MaxMana, ColorReset)
	fmt.Printf("  Or     : %s%d pièces%s\n", ColorYellow, c.Money, ColorReset)
	fmt.Println("--------------------------------------")
	fmt.Println("  Équipement :")
	if c.Equipment.Tete != "" {
		fmt.Printf("    Tête  : %s\n", c.Equipment.Tete)
	} else {
		fmt.Println("    Tête  : (vide)")
	}
	if c.Equipment.Torse != "" {
		fmt.Printf("    Torse : %s\n", c.Equipment.Torse)
	} else {
		fmt.Println("    Torse : (vide)")
	}
	if c.Equipment.Pieds != "" {
		fmt.Printf("    Pieds : %s\n", c.Equipment.Pieds)
	} else {
		fmt.Println("    Pieds : (vide)")
	}
	fmt.Println("--------------------------------------")
	fmt.Println("  Sorts connus :")
	for _, skill := range c.Skills {
		fmt.Printf("    - %s\n", skill)
	}
	fmt.Println("======================================")
}

// isDead vérifie si les PV du personnage sont à 0 ou moins, et gère une mécanique de "résurrection".
func isDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Printf(ColorRed+`
      _____
     /     \
    | () () |
     \  ^  /
      |||||
    💀 %s est tombé à 0 PV !
`+ColorReset+"\n", c.Name)
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf(ColorGreen+"✨ %s revient à la vie avec %d PV.\n\n"+ColorReset, c.Name, c.CurrentHP)
		return true
	}
	return false
}

// gainXP ajoute de l'expérience et gère la montée de niveau (et l'augmentation des stats) si le palier est atteint.
func gainXP(c *Character, amount int) {
	c.CurrentXP += amount
	fmt.Printf(ColorCyan+"✨ Vous gagnez %d points d'expérience !\n"+ColorReset, amount)

	for c.CurrentXP >= c.MaxXP {
		c.Level++
		c.CurrentXP -= c.MaxXP
		c.MaxXP = int(float64(c.MaxXP) * 1.5)
		c.MaxHP += 10
		c.CurrentHP = c.MaxHP
		fmt.Printf(ColorGreen+`
       /\
      /__\   🆙 NIVEAU SUPÉRIEUR !
     /\  /\     Vous passez niveau %d !
    /__\/__\    Vos PV Max augmentent de 10. PV restaurés.
`+ColorReset+"\n\n", c.Level)
	}
}

// hasSkill est une fonction utilitaire pour vérifier si le joueur connaît déjà un sort donné.
func hasSkill(c *Character, skill string) bool {
	for _, s := range c.Skills {
		if s == skill {
			return true
		}
	}
	return false
}
