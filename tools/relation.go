package tools

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
)

func GetRelation() models.Index {
	var relation models.Index
	var w http.ResponseWriter

	path := GetAPI("https://groupietrackers.herokuapp.com/api")

	result, err := http.Get(path["relation"])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		MessageError(err)
	}
	json.NewDecoder(result.Body).Decode(&relation)
	return relation
}

func ReturnRelations(id int) map[string][]string {
	var relations models.Relation
	lien := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	r, err := http.Get(lien)
	if err != nil {
		var w http.ResponseWriter
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		MessageError(err)
	}

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&relations)
	if err != nil {
		var w http.ResponseWriter
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, the server encountered an unexpected error.")
		MessageError(err)
	}
	return relations.DatesLocations
}
