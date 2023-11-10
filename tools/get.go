package tools

import (
	"encoding/json"
	"groupie/models"
	"net/http"
	"sync"
)

func GetAPI(URL string) (paths map[string]string) {
	response, err := http.Get(URL)
	MessageError(err)
	MessageError(json.NewDecoder(response.Body).Decode(&paths))
	return
}

func GetData(wg *sync.WaitGroup) {
	defer wg.Done() // Marque la fin de l'exécution de la goroutine
	models.All_Artists = GetArtist()
	models.All_Locations = GetLocation()
}
