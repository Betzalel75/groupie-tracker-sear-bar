package tools

import (
	"net/http"
	"text/template"
)

func HandleArtist(w http.ResponseWriter, r *http.Request) {
	artists := GetArtist()
	t, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server", "Sorry, the server not match.")
		return
	}
	allLocation := GetLocation()
	data := Data{
		AllArtists: artists,
		AllLocations: allLocation,
	}
	t.Execute(w, data)

}
