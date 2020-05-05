package response

import (
	"encoding/json"
	"html/template"
	"net/http"
)

//JSON return json response
func JSON(w http.ResponseWriter, m interface{}, s int) {
	data, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	w.Write(data)
}

//Render renders html to client
func Render(w http.ResponseWriter, data interface{}, filenames ...string) {
	tem, err := template.ParseFiles(filenames...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tem.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
