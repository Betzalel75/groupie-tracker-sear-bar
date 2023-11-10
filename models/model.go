package models

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
	Id             int
	DatesLocations map[string][]string
}

type Data struct {
	Artist         Artist
	Artists        []Artist
	AllArtists     []Artist
	SingleLocation []string
	AllLocations   [][]string
	Date           Date
	Relations      Index
	Locations      map[int]map[string][]string
	UniqueMembers []string // Ajoutez un champ pour stocker les membres uniques
	UniqueLocations []string // Ajoutez un champ pour stocker les localisations uniques
	// ArtistInfo []struct{Name string; CreationDate int}
	ArtistInfos []ArtistInfo

}

type ArtistInfo struct{
	Name string
	CreationDate int
}

type Index struct {
	Index []Relation
}

var All_Artists []Artist
var All_Locations [][]string
