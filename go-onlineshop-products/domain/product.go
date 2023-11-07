package domain

import "encoding/json"

type Product struct {
	Id         int
	CategoryId int
	Name       string
	Discount   int
	Components []Product
}

func (b Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}
