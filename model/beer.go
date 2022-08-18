package model

type Beer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}
