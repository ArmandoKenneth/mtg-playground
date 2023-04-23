package main

import (
	"fmt"
	"log"

	"github.com/ArmandoKenneth/mtg-playground/services/apis/scryfall"
)

func main() {
	// 2023/04/23 13:53:13 G,U
	// 2023/04/23 13:53:13 Mutagen Connoisseur
	// 2023/04/23 13:53:13 4a84b6d7-3944-4223-ac16-aa5e59ac84cb
	card, error := scryfall.GetCard("4a84b6d7-3944-4223-ac16-aa5e59ac84cb")
	// card, error := scryfall.GetCard("")
	if error != nil {
		fmt.Println(error)
	} else {
		log.Printf(card.Name)
	}
}
