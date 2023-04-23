package scryfall

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func makeScryfallRequest(path string) ([]byte, error) {
	response, error := http.Get("https://api.scryfall.com/" + path)
	if error != nil {
		log.Fatalln(error)
	}
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		var errorResponse []byte
		return errorResponse, errors.New("Error making request")
	}

	return body, nil
}
