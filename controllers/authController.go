package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"newproject/database"
	"newproject/helpers"
	"newproject/models"
	"newproject/response"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type access struct {
	Token string `json:"token"`
}

// Register handler
func Register(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := models.User{}
	user, err := u.Create(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSON(w, user, http.StatusOK)
}

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cred := credentials{}
	err = json.Unmarshal(b, &cred)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.DB
	u := models.User{}
	db.Where("Username = ?", cred.Username).First(&u)

	if u.Username != cred.Username || !helpers.CheckPasswordHash(cred.Password, string(u.Password)) {
		http.Error(w, "Wrong credentials", http.StatusUnauthorized)
		return
	}

	a := access{}
	a.Token, err = helpers.ObtainToken(u.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, a, http.StatusOK)
}
