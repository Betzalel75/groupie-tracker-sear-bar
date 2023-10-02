package tools

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, "Page Not Found", "Sorry, the page you are looking for could not be found.")
		return
	}
	// Générer une nouvelle source avec une valeur de graine spécifique
	source := rand.NewSource(time.Now().UnixNano())

	// Créer un nouveau générateur aléatoire en utilisant la source
	random := rand.New(source)

	// Générer un nombre aléatoire entre 1 et 52
	randomNumber := random.Intn(52) + 1

	artist, err := ReturnArtist(randomNumber)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}
	allArtists := GetArtist()
	allLocations := GetLocation()
	data := Data{
		Artist:       artist,
		AllArtists:   allArtists,
		AllLocations: allLocations,
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		return
	}
	t.Execute(w, data)
}
