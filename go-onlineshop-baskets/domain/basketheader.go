package domain

import (
	"encoding/json"

	. "github.com/ahmetb/go-linq/v3"
)

type BasketHeader struct {
	PromoCode         string        `json:"promoCode"`
	Note              string        `json:"note"`
	ShippingAddressId string        `json:"shippingAddressId"`
	BillingAddressId  string        `json:"billingAddressId"`
	Items             *[]BasketItem `json:"items"`
}

func (b BasketHeader) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b BasketHeader) HasSameItem(item BasketItem) bool {
	return From(*b.Items).AnyWith(func(c interface{}) bool {
		return c.(BasketItem).Equals(item)
	})
}
