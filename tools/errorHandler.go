package tools

import (
	"net/http"
	"text/template"
)

type Error struct {
	Code    int
	Warning string
	Message string
}

func ErrorHandler(w http.ResponseWriter, code int, warning string, message string) {
	error := Error{
		Code:    code,
		Warning: warning,
		Message: message,
	}
	w.WriteHeader(error.Code)
	t, err := template.ParseFiles("./templates/errors.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "errors", error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
