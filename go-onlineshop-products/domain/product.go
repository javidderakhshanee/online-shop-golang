package domain

import "encoding/json"

type Product struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Discount   int    `json:"discount"`
}

func (b Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}
