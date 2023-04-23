package scryfall_interfaces

type Card struct {
	Id            string   `json:"id"`
	ColorIdentity []string `json:"color_identity"`
	Name          string   `json:"name"`
}
