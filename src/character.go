package main

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Character struct {
	Name         string
	Class        string
	MaxHP        int
	CurrentHP    int
	Level        int
	CurrentXP    int
	MaxXP        int
	Inventory    []string
	InventoryMax int
	Money        int
	Skills       []string
	Equipment    Equipment
}

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
	case "Elfe":
		c.MaxHP = 80
		c.CurrentHP = 40
	case "Nain":
		c.MaxHP = 120
		c.CurrentHP = 60
	default:
		c.MaxHP = 100
		c.CurrentHP = 50
	}

	return c
}

func displayInfo(c *Character) {
	fmt.Println("========== FICHE PERSONNAGE ==========")
	fmt.Printf("  Nom    : %s\n", c.Name)
	fmt.Printf("  Classe : %s\n", c.Class)
	fmt.Printf("  Niveau : %d (XP : %d / %d)\n", c.Level, c.CurrentXP, c.MaxXP)
	fmt.Printf("  PV     : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("  Or     : %d pièces\n", c.Money)
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

func isDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Printf("\n💀 %s est tombé à 0 PV !\n", c.Name)
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("✨ %s revient à la vie avec %d PV.\n\n", c.Name, c.CurrentHP)
		return true
	}
	return false
}

func gainXP(c *Character, amount int) {
	c.CurrentXP += amount
	fmt.Printf("✨ Vous gagnez %d points d'expérience !\n", amount)

	for c.CurrentXP >= c.MaxXP {
		c.Level++
		c.CurrentXP -= c.MaxXP
		c.MaxXP = int(float64(c.MaxXP) * 1.5)
		c.MaxHP += 10
		c.CurrentHP = c.MaxHP
		fmt.Printf("\n🆙 NIVEAU SUPÉRIEUR ! Vous passez niveau %d !\n   Vos PV Max augmentent de 10. PV restaurés.\n\n", c.Level)
	}
}

func hasSkill(c *Character, skill string) bool {
	for _, s := range c.Skills {
		if s == skill {
			return true
		}
	}
	return false
}