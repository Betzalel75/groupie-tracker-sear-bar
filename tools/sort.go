package tools

import (
	"groupie/models"
)

func Sort() []string {
	uniqueMembers := make(map[string]struct{}) // Utilisez une carte pour garder les membres uniques

	for i := 0; i < len(models.All_Artists); i++ {
		for j := 0; j < len(models.All_Artists[i].Members); j++ {
				uniqueMembers[models.All_Artists[i].Members[j]] = struct{}{}
		}
	}
	// Convertir la carte en une liste de membres uniques
	tabUnique := make([]string, 0, len(uniqueMembers))
	for member := range uniqueMembers {
		tabUnique = append(tabUnique, member)
	}
	/*for i := 0; i < len(models.All_Artists); i++ {
		for j := 0; j < len(tabUnique); j++ {
			if models.All_Artists[i].Name == tabUnique[j] {
				tabUnique[j] = ""
			}
		}
	}*/
	return tabUnique
}

func SortLocation() []string {
	uniqueLocations := make(map[string]struct{}) // Utilisez une carte pour garder les localisations uniques

	for i := 0; i < len(models.All_Locations); i++ {
		for j := 0; j < len(models.All_Locations[i]); j++ {
			uniqueLocations[models.All_Locations[i][j]] = struct{}{}
		}
	}
	// Convertir la carte en une liste de localisations uniques
	uniqueLocationList := make([]string, 0, len(uniqueLocations))
	for location := range uniqueLocations {
		uniqueLocationList = append(uniqueLocationList, location)
	}

	return uniqueLocationList
}

// Tableau des dates associer au nom de l'artist

func TabDates_Artists() []models.ArtistInfo {
  
    // Créez un nouveau tableau pour stocker les informations de chaque artiste
    var ArtistInfo []models.ArtistInfo

    for _, artist := range models.All_Artists {
        // Ajoutez les informations requises au nouveau tableau
        artistInfo := models.ArtistInfo {
            Name:         artist.Name,
            CreationDate: artist.CreationDate,
        }
        ArtistInfo = append(ArtistInfo, artistInfo)
    }
    return ArtistInfo
}