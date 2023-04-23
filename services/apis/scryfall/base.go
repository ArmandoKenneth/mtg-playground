package scryfall

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	scryfall_interfaces "github.com/ArmandoKenneth/mtg-playground/services/apis/scryfall/interfaces"
)

func makeScryfallRequest(path string) []byte {
	response, error := http.Get("https://api.scryfall.com/" + path)
	if error != nil {
		log.Fatalln(error)
	}
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatalln(error)
		// raise error
	}

	return body
}

func GetCard(cardId string) scryfall_interfaces.Card {
	path := "cards/" + cardId
	body := makeScryfallRequest(path)

	var card scryfall_interfaces.Card
	error := json.Unmarshal(body, &card)
	if error != nil {

	}
	log.Printf(strings.Join(card.ColorIdentity, ","))
	log.Printf(card.Name)
	log.Printf(card.Id)

	return card
}
