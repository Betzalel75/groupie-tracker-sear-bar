package tools

import (
	"net/http"
)

// Vérifie la disponibilité d'Internet en effectuant une requête de ping simple
func IsInternetAvailable() bool {
	_, err := http.Get("https://www.google.com")
	if err != nil {
		return false
	}
	return true
}
