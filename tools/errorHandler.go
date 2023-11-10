package tools

import (
	"net/http"
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
	RenderTemplate(w, "errors", "error", error)
}
