package models

import "github.com/jinzhu/gorm"

//Order defines model for order
type Order struct {
	gorm.Model
	CustomerID        uint    `json:"customerID"`
	NumItemsSold      int     `json:"numItemsSold"`
	ShippingAddress   string  `json:"shippingAddress"`
	ShippingInterior  string  `json:"shippingInterior"`
	ShippingCity      string  `json:"shippingCity"`
	ShippingCountry   string  `json:"shippingCountry"`
	BillingAddress    string  `json:"billingAddress"`
	BillingInterior   string  `json:"billingInterior"`
	BillingCity       string  `json:"billingCity"`
	BillingCountry    string  `json:"billingCountry"`
	TotalAmount       float32 `json:"totalAmount"`
	ShippingTotal     float32 `json:"shippingTotal"`
	NetTotal          float32 `json:"netTotal"`
	PaymentMethodID   uint    `json:"paymentMethodID"`
	ReturningCustomer bool    `json:"returningCustomer"`
	StatusID          uint    `json:"statusID"`
}

//OrderItem defines Order Item
type OrderItem struct {
	gorm.Model
	OrderID        uint    `json:"orderID"`
	ProductID      uint    `json:"productID"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	ProductRevenue float32 `json:"productRevenue"`
}

// Status defines model for statuses
type Status struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(100);unique_index"`
}
