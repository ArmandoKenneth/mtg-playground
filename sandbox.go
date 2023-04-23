package main

import (
	"github.com/ArmandoKenneth/mtg-playground/services/apis/scryfall"
)

func main() {
	// var result bool = colors.MatchesColorIdentify(1, 2)
	// var resultText string
	// if result {
	// 	resultText = "Yes"
	// } else {
	// 	resultText = "No"
	// }
	// fmt.Printf("Matches color identity? %s\n", resultText)

	// 2023/04/23 13:53:13 G,U
	// 2023/04/23 13:53:13 Mutagen Connoisseur
	// 2023/04/23 13:53:13 4a84b6d7-3944-4223-ac16-aa5e59ac84cb
	scryfall.GetCard("4a84b6d7-3944-4223-ac16-aa5e59ac84cb")
}
