package handlers

import (
	"groupie/models"
	"groupie/tools"
	"net/http"
)

func HandleArtist(w http.ResponseWriter, r *http.Request, artists []models.Artist, allLocation [][]string) {
	data := models.Data{
		AllArtists:   artists,
		AllLocations: allLocation,
	}
	data.UniqueMembers = tools.Sort()
	data.UniqueLocations = tools.SortLocation()
	data.ArtistInfos = tools.TabDates_Artists()
	tools.RenderTemplate(w, "artist", "artist", data)
}
