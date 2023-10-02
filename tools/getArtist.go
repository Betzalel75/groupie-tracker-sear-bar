package tools

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	Id           int
	Name         string
	Members      []string
	Image        string
	CreationDate int
	FirstAlbum   string
	Location     []string
	ConcertDate  string
	Relations    string
}

type Location struct {
	Id        int
	Locations []string
	Dates     string
}

type Date struct {
	Id    int
	Dates []string
}

type Relation struct {
	Id        int
	Relations map[string]string
}

type Data struct {
	Artist         Artist
	Artists        []Artist
	AllArtists     []Artist
	SingleLocation []string
	AllLocations   [][]string
	Date           Date
}

func GetArtist() []Artist {
	var w http.ResponseWriter
	var artists []Artist
	lien := "https://groupietrackers.herokuapp.com/api/artists"
	r, err := http.Get(lien)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		return nil
	}
	json.NewDecoder(r.Body).Decode(&artists)
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
	//fmt.Println(location)
	return location
}
