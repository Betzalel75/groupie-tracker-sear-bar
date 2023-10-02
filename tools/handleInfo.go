package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ReturnArtist(id int) (Artist, error) {
	var artist Artist
	lien := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	r, err := http.Get(lien)
	if err != nil {
		return artist, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

func ReturnLocations(id int) ([]string, error) {
	var location Location
	lien := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)
	r, err := http.Get(lien)
	if err != nil {
		return location.Locations, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		return location.Locations, err
	}
	return location.Locations, nil
}

func ReturnDates(id int) (Date, error) {
	var dates Date
	lien := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	r, err := http.Get(lien)
	if err != nil {
		return dates, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&dates)
	if err != nil {
		return dates, err
	}
	tabDates := dates.Dates // Copier le contenu de tout le tableau des dates
	dates.Dates = nil       // Vider le tableau des dates
	for _, v := range tabDates {
		v = strings.TrimPrefix(v, "*")
		dates.Dates = append(dates.Dates, v)
	}
	return dates, nil
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	idArtistStr := r.URL.Query().Get("id")
	idArtist, err := strconv.Atoi(idArtistStr)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, "Bad Request", "You've sent a Bad Request.")
		return
	}

	if idArtist <= 0 {
		ErrorHandler(w, http.StatusBadRequest, "Bad Request", "You've sent a Bad Request.")
		return
	}
	if idArtist > 52 {
		ErrorHandler(w, http.StatusNotFound, "Page Not Found", "Sorry, the page you are looking for could not be found.")
		return
	}

	artist, err := ReturnArtist(idArtist)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}
	location, err := ReturnLocations(idArtist)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}
	dates, err := ReturnDates(idArtist)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}
	allArtists := GetArtist()
	allLocation := GetLocation()

	data := Data{
		Artist:         artist,
		AllArtists:     allArtists,
		SingleLocation: location,
		AllLocations:   allLocation,
		Date:           dates,
	}
	t, err := template.ParseFiles("templates/info.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		return
	}
	t.Execute(w, data)
}
