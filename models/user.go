package models

import (
	"encoding/json"
	"newproject/database"
	"newproject/helpers"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password"`
	IsStaff   bool   `json:"isStaff"`
	IsAdmin   bool   `json:"isAdmin"`
}

//UserSerializer implementation
type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	IsStaff   bool   `json:"isStaff"`
	IsAdmin   bool   `json:"isAdmin"`
}

func serializeUser(u User) UserSerializer {
	return UserSerializer{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		IsStaff:   u.IsStaff,
		IsAdmin:   u.IsAdmin,
	}
}

//MarshalJSON interface implementation
func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(serializeUser(u))
}

//All instances
func (u User) All() interface{} {
	db := database.DB
	users := []User{}
	db.Find(&users)
	return users
}

//Create instance
func (u User) Create(data []byte) (interface{}, error) {
	db := database.DB

	err := json.Unmarshal(data, &u)

	if err != nil {
		return nil, err
	}

	pwd := u.Password
	pwd, err = helpers.HashPassword(pwd)
	if err != nil {
		return nil, err
	}

	u.Password = pwd

	db.Create(&u)
	return u, nil
}

//Find instance
func (u User) Find(id int) interface{} {
	db := database.DB
	db.First(&u, id)
	return u
}

//Update instance
func (u User) Update(id int, data []byte) (interface{}, error) {
	db := database.DB
	update := User{}
	err := json.Unmarshal(data, &update)
	if err != nil {
		return nil, err
	}
	db.First(&u, id).Updates(update)
	return u, nil
}

//Delete instance
func (u User) Delete(id uint) {
	db := database.DB
	u.ID = id
	db.Delete(&u)
}
