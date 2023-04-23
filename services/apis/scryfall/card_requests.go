package scryfall

import (
	"encoding/json"
	"errors"

	scryfall_interfaces "github.com/ArmandoKenneth/mtg-playground/services/apis/scryfall/interfaces"
)

func GetCard(cardId string) (scryfall_interfaces.Card, error) {
	var card scryfall_interfaces.Card

	if cardId == "" {
		return card, errors.New("Invalid cardId")
	}

	path := "cards/" + cardId
	body, error := makeScryfallRequest(path)

	// NEXT: learn about error handling properly
	if error != nil {
		return card, error
	}

	error = json.Unmarshal(body, &card)
	if error != nil {
		return card, errors.New("Bad data")
	}

	return card, nil
}
