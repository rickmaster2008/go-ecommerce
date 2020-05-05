package controllers

import (
	"io/ioutil"
	"net/http"
	"newproject/models"
	"newproject/response"
	"strconv"

	"github.com/gorilla/mux"
)

//Controller defines controller interface
type Controller interface {
	Index(http.ResponseWriter,
		*http.Request)
	Store(http.ResponseWriter,
		*http.Request)
	Show(http.ResponseWriter,
		*http.Request)
	Update(http.ResponseWriter,
		*http.Request)
	Destroy(http.ResponseWriter,
		*http.Request)
}

//BaseController defines base controller
type BaseController struct {
	M models.Model
}

// Index Lists all
func (bc BaseController) Index(w http.ResponseWriter,
	r *http.Request) {
	m := bc.M
	ms := m.All()
	response.JSON(w, ms, http.StatusOK)
}

//Store store one
func (bc BaseController) Store(w http.ResponseWriter,
	r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := bc.M.Create(b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, m, http.StatusCreated)
}

//Show retrieves one
func (bc BaseController) Show(w http.ResponseWriter,
	r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := bc.M.Find(id)

	response.JSON(w, m, http.StatusOK)
}

//Update updates one
func (bc BaseController) Update(w http.ResponseWriter,
	r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := bc.M.Update(id, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, m, http.StatusOK)
}

//Destroy deletes one
func (bc BaseController) Destroy(w http.ResponseWriter,
	r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bc.M.Delete(uint(id))
	w.WriteHeader(http.StatusOK)
}
