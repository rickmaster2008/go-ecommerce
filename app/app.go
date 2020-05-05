package app

import (
	"fmt"
	"log"
	"net/http"
	"newproject/controllers"

	"github.com/gorilla/mux"
)

//App defines new http app
type App struct {
	Router *mux.Router
}

//NewApp returns an instance of App
func NewApp() App {
	return App{Router: mux.NewRouter()}
}

//Static defines am static file server
func (app *App) Static(dirPath string, staticURL string) {
	fs := http.FileServer(http.Dir(dirPath))
	app.Router.PathPrefix(staticURL).Handler(http.StripPrefix(staticURL, fs))
}

//Resource defines a CRUD view set
func (app *App) Resource(path string, c controllers.Controller, mws ...func(http.Handler) http.Handler) {
	s := app.Router.PathPrefix(path).Subrouter()
	s.Methods(http.MethodGet).HandlerFunc(c.Index)
	s.Methods(http.MethodPost).HandlerFunc(c.Store)
	s.Path("/{id:[0-9]+}").Methods(http.MethodGet).HandlerFunc(c.Show)
	s.Path("/{id:[0-9]+}").Methods(http.MethodPut).HandlerFunc(c.Update)
	s.Path("/{id:[0-9]+}").Methods(http.MethodDelete).HandlerFunc(c.Destroy)
	for _, mw := range mws {
		s.Use(mw)
	}
}

//Get defines a function to habdle GET Method
func (app *App) Get(path string, f func(http.ResponseWriter, *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

//Post defines a function to habdle POST Method
func (app *App) Post(path string, f func(http.ResponseWriter, *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

//Put defines a function to habdle PUT Method
func (app *App) Put(path string, f func(http.ResponseWriter, *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

//Delete defines a function to habdle DELETE Method
func (app *App) Delete(path string, f func(http.ResponseWriter, *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

//Listen defines function for serve and listen
func (app *App) Listen(port string) {
	http.Handle("/", app.Router)
	fmt.Println("server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func parserMiddleare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
