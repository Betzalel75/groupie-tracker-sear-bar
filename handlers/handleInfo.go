package handlers

import (
	"groupie/models"
	"groupie/tools"
	"net/http"
	"strconv"
)

var APIURL = "https://groupietrackers.herokuapp.com/api"

func HandleInfo(w http.ResponseWriter, r *http.Request, allArtists []models.Artist, allLocation [][]string) {
	idArtistStr := r.URL.Query().Get("id")
	idArtist, err := strconv.Atoi(idArtistStr)
	if err != nil {
		tools.ErrorHandler(w, http.StatusBadRequest, "Bad Request", "You've sent a Bad Request.")
		return
	}

	if idArtist <= 0 {
		tools.ErrorHandler(w, http.StatusBadRequest, "Bad Request", "You've sent a Bad Request.")
		return
	}
	if idArtist > 52 {
		tools.ErrorHandler(w, http.StatusNotFound, "Page Not Found", "Sorry, the page you are looking for could not be found.")
		return
	}

	artist, err := tools.ReturnArtist(idArtist)
	if err != nil {
		tools.ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}
	locations, err := tools.ReturnLocations(idArtist)
	if err != nil {
		tools.ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}
	dates, err := tools.ReturnDates(idArtist)
	if err != nil {
		tools.ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}

	// Création d'une map pour stocker les emplacements pour chaque ID de relation
	locationsRelation := make(map[int]map[string][]string)

	// Appel à la fonction ReturnRelations pour obtenir les emplacements associés à l'ID de relation actuel
	locationsRelation[idArtist] = tools.ReturnRelations(idArtist)

	data := models.Data{
		Artist:         artist,
		AllArtists:     allArtists,
		SingleLocation: locations,
		AllLocations:   allLocation,
		Date:           dates,
		Locations:      locationsRelation,
	}
	data.UniqueMembers = tools.Sort()
	data.UniqueLocations = tools.SortLocation()
	data.ArtistInfos = tools.TabDates_Artists()
	tools.RenderTemplate(w, "info", "info", data)
}
