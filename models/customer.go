package models

import "github.com/jinzhu/gorm"

// Customer defines customer model
type Customer struct {
	gorm.Model
	UserID           uint   `json:"userID"`
	ShippingAddress  string `json:"shippingAddress"`
	ShippingInterior string `json:"shippingInterior"`
	ShippingCity     string `json:"shippingCity"`
	ShippingCountry  string `json:"shippingCountry"`
	BillingAddress   string `json:"billingAddress"`
	BillingInterior  string `json:"billingInterior"`
	BillingCity      string `json:"billingCity"`
	BillingCountry   string `json:"billingCountry"`
}
