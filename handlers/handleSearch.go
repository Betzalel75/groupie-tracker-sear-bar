package handlers

import (
	"groupie/models"
	"groupie/tools"
	"net/http"
	"strconv"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request, allArtists []models.Artist, All_Locations [][]string) {
	query := r.URL.Query().Get("search") // Obtient la valeur de la barre de recherche
	date, _ := strconv.Atoi(query)
	query = strings.ToLower(query)

	for i := 0; i < len(allArtists); i++ {
		allArtists[i].Location = All_Locations[i]
	}
	var matchingArtists []models.Artist // Liste pour stocker les artistes correspondants
	for i := 0; i < len(allArtists); i++ {
		lowerName := strings.ToLower(allArtists[i].Name) // Convertir le nom de l'artiste en minuscules

		if strings.Contains(lowerName, query) || strings.EqualFold(allArtists[i].FirstAlbum, query) ||
			date == allArtists[i].CreationDate {
			matchingArtists = append(matchingArtists, allArtists[i])
			continue // Passer au prochain artiste après avoir trouvé une correspondance
		}
		for j := 0; j < len(allArtists[i].Members); j++ {
			lowerMember := strings.ToLower(allArtists[i].Members[j])
			if strings.Contains(lowerMember, query) {
				matchingArtists = append(matchingArtists, allArtists[i])
				break // Passer au prochain artiste après avoir trouvé une correspondance avec un membre
			}
		}
		artistLocations := All_Locations[allArtists[i].Id-1]
		for _, loc := range artistLocations {
			lowerLocation := strings.ToLower(loc)
			if strings.Contains(lowerLocation, query) {
				matchingArtists = append(matchingArtists, allArtists[i])
				break // Passer au prochain artiste après avoir trouvé une correspondance avec une location
			}
		}
	}

	data := models.Data{
		Artists:      matchingArtists,
		AllArtists:   allArtists,
		AllLocations: All_Locations,
	}

	data.UniqueMembers = tools.Sort()
	data.UniqueLocations = tools.SortLocation()
	data.ArtistInfos = tools.TabDates_Artists()
	tools.RenderTemplate(w, "results", "results", data)
}
