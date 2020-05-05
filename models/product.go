package models

import (
	"encoding/json"
	"newproject/database"

	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	gorm.Model
	SKU           string  `json:"sku" gorm:"type:varchar(20);unique_index"`
	CategoryID    uint    `json:"categoryID"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float32 `json:"price"`
	DiscountPrice float32 `json:"discountPrice"`
	Inventory     int     `json:"inventory"`
	InInventory   bool    `json:"inInventory"`
}

// Category defines product categories
type Category struct {
	gorm.Model
	Name   string `json:"name"`
	Parent uint   `json:"parent"`
}

//Delete instance
func (c Category) Delete(id uint) {
	db := database.DB
	c.ID = id
	db.Delete(&c)
}

//All instances
func (p Product) All() interface{} {
	db := database.DB
	products := []Product{}
	db.Find(&products)
	return products
}

//Create instance
func (p Product) Create(data []byte) (interface{}, error) {
	db := database.DB
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	db.Create(&p)
	return p, nil
}

//Find instance
func (p Product) Find(id int) interface{} {
	db := database.DB
	db.First(&p, id)
	return p
}

//Update instance
func (p Product) Update(id int, data []byte) (interface{}, error) {
	db := database.DB
	update := Product{}
	err := json.Unmarshal(data, &update)
	if err != nil {
		return nil, err
	}
	db.First(&p, id).Updates(update)
	return p, nil
}

//Delete instance
func (p Product) Delete(id uint) {
	db := database.DB
	p.ID = id
	db.Delete(&p)
}

// All Category instances
func (c Category) All() interface{} {
	db := database.DB
	cats := []Category{}
	db.Find(&cats)
	return cats
}

//Create instance
func (c Category) Create(data []byte) (interface{}, error) {
	db := database.DB
	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	db.Create(&c)
	return c, nil
}

//Find instance
func (c Category) Find(id int) interface{} {
	db := database.DB
	db.First(&c, id)
	return c
}

//Update instance
func (c Category) Update(id int, data []byte) (interface{}, error) {
	db := database.DB
	update := Category{}
	err := json.Unmarshal(data, &update)
	if err != nil {
		return nil, err
	}
	db.First(&c, id).Updates(update)
	return c, nil
}
