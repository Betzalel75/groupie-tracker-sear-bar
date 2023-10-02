package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search") // Obtient la valeur de la barre de recherche
	date, _ := strconv.Atoi(query)
	artists, err := SearchArtist(query)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artists.")
		return
	}
	nbr := len(artists)
	i := 1
	for nbr != 1 {
		lc, _ := ReturnLocations(i)
		artists[i].Location = lc
		i++
		nbr -= 1
	}
	var matchingArtists []Artist // Liste pour stocker les artistes correspondants
	alllocations := GetLocation()
	for i := 0; i < len(artists); i++ {
		lowerName := strings.ToLower(artists[i].Name) // Convertir le nom de l'artiste en minuscules
		if strings.EqualFold(query, lowerName) || strings.EqualFold(query, artists[i].FirstAlbum) ||
			date == artists[i].CreationDate {
			matchingArtists = append(matchingArtists, artists[i])
			continue // Passer au prochain artiste après avoir trouvé une correspondance
		}
		for j := 0; j < len(artists[i].Members); j++ {
			lowerMember := strings.ToLower(artists[i].Members[j])
			if strings.EqualFold(query, lowerMember) {
				matchingArtists = append(matchingArtists, artists[i])
				break // Passer au prochain artiste après avoir trouvé une correspondance avec un membre
			}
		}
		artistLocations := alllocations[artists[i].Id-1]
		for _, loc := range artistLocations {
			lowerLocation := strings.ToLower(loc)
			if strings.EqualFold(query, lowerLocation) {
				matchingArtists = append(matchingArtists, artists[i])
				break // Passer au prochain artiste après avoir trouvé une correspondance avec une location
			}
		}
	}

	allArtists := GetArtist()
	data := Data{
		Artists:      matchingArtists,
		AllArtists:   allArtists,
		AllLocations: alllocations,
	}

	t, err := template.ParseFiles("templates/results.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		return
	}

	t.Execute(w, data)
}

func SearchArtist(query string) ([]Artist, error) {
	var artists []Artist
	lien := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists?q=%s", url.QueryEscape(query))
	r, err := http.Get(lien)
	if err != nil {
		return artists, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&artists)
	if err != nil {
		return artists, err
	}
	return artists, nil
}
