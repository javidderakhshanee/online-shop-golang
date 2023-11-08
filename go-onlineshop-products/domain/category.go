package domain

import "encoding/json"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (b Category) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}
