package main

import (
	"fmt"

	logic "projet-red/src"
)

func main() {
	logic.ClearScreen()
	logic.PrintBanner()
	logic.WaitForInput()

	player := logic.Introduction()
	logic.Quest1_TheVillage(&player)
	logic.Quest2_TheForest(&player)
	logic.Quest3_TheFortress(&player)

	fmt.Println(logic.ColorYellow + "\nÀ suivre dans le prochain chapitre..." + logic.ColorReset)
	fmt.Println("Merci d'avoir joué à la démo de Legends of Ynov !")
}
