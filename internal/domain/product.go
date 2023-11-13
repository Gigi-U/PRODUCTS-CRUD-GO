package domain

type Product struct {
	ID       int     `json:"id"`
	Name   string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price   float64 `json:"price"`
 }