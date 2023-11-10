package tools

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
	"strings"
)

var APIURL = "https://groupietrackers.herokuapp.com/api"

func GetArtist() []models.Artist {
	var w http.ResponseWriter
	var artists []models.Artist

	path := GetAPI("https://groupietrackers.herokuapp.com/api")

	result, err := http.Get(path["artists"])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		MessageError(err)
		return nil
	}
	json.NewDecoder(result.Body).Decode(&artists)
	return artists
}

func GetLocation() [][]string {
	var w http.ResponseWriter
	var location [][]string
	nbr := len(GetArtist())
	for i := 1; i <= nbr; i++ {
		locations, err := ReturnLocations(i)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
			return nil
		}
		location = append(location, locations)
	}
	return location
}

func ReturnLocations(id int) ([]string, error) {
	var location models.Location
	lien := fmt.Sprintf("%s/locations/%d", APIURL, id)
	r, err := http.Get(lien)
	if err != nil {
		return location.Locations, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		return location.Locations, err
	}
	tabDates := location.Locations // Copier le contenu de tout le tableau des localisation
	location.Locations = nil       // Vider le tableau des localisations
	for _, v := range tabDates {
		v = strings.ReplaceAll(v, "_", "-")
		location.Locations = append(location.Locations, v)
	}
	
	return location.Locations, nil
}

func ReturnArtist(id int) (models.Artist, error) {
	var artist models.Artist

	for i := 0; i < len(models.All_Artists); i++ {
		if i == id-1 {
			artist = models.All_Artists[i]
		}
	}
	return artist, nil
}

func ReturnDates(id int) (models.Date, error) {
	var dates models.Date
	lien := fmt.Sprintf("%s/dates/%d", APIURL, id)
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
