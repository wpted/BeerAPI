package model

import (
	"encoding/json"
	"log"
)

// Beer is a model that matches the field within the PSQL database
type Beer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}

// MarshalBinary marshals the beer struct and return a []byte and an error
func (b Beer) MarshalBinary() ([]byte, error) {
	json, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err.Error())
	}
	return json, err
}
