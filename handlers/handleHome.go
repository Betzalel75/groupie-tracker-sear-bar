package handlers

import (
	"groupie/models"
	"groupie/tools"
	"math/rand"
	"net/http"
	"time"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tools.ErrorHandler(w, http.StatusNotFound, "Page Not Found", "Sorry, the page you are looking for could not be found.")
		return
	}
	// Générer une nouvelle source avec une valeur de graine spécifique
	source := rand.NewSource(time.Now().UnixNano())

	// Créer un nouveau générateur aléatoire en utilisant la source
	random := rand.New(source)

	// Générer un nombre aléatoire entre 1 et 52
	randomNumber := random.Intn(52) + 1

	artist := models.All_Artists[randomNumber]

	data := models.Data{
		Artist:     artist,
		AllArtists: models.All_Artists,
	}
	data.UniqueMembers = tools.Sort()
	data.UniqueLocations = tools.SortLocation()
	data.ArtistInfos = tools.TabDates_Artists()
	tools.RenderTemplate(w, "index", "index", data)
}
