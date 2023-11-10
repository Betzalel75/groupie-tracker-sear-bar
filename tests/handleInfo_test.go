package main

import (
	"encoding/json"
	"fmt"
	"groupie/tools"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Location struct {
	Locations []string `json:"locations"`
}
type Date struct {
	Id    int `json:"id"`
	Dates []string `json:"dates"`
}

func TestReturnLocations(t *testing.T) {
	// Créer un serveur de test mock
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Location{
			Locations: []string{"Location1", "Location2", "Location3"},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Remplacer l'URL de l'API dans le package tools
	tools.APIURL = mockServer.URL

	// Appeler la fonction à tester
	locations, err := tools.ReturnLocations(1)

	// Vérifier les résultats
	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}

	expectedLocations := []string{"Location1", "Location2", "Location3"}
	if len(locations) != len(expectedLocations) {
		t.Errorf("Le nombre de localisations retourné (%d) ne correspond pas au nombre attendu (%d)", len(locations), len(expectedLocations))
	}

	for i, loc := range locations {
		fmt.Println(loc)
		if loc != expectedLocations[i] {
			t.Errorf("Localisation différente à l'index %d, attendu : %s, obtenu : %s", i, expectedLocations[i], loc)
		}
	}
}

func TestReturnDates(t *testing.T) {
	// Créer un serveur de test mock
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Date{
			Dates: []string{"Date1", "Date2", "Date3", "Date4"},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Remplacer l'URL de l'API dans le package tools
	tools.APIURL = mockServer.URL

	// Appeler la fonction à tester
	date, err := tools.ReturnDates(12)
	dates := date.Dates
	// Vérifier les résultats
	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}

	expectedDates := []string{"Date1", "Date2", "Date3", "Date4"}
	if len(dates) != len(expectedDates) {
		t.Errorf("Le nombre de Dates retourné (%d) ne correspond pas au nombre attendu (%d)", len(dates), len(expectedDates))
	}

	for i, loc := range dates {
		if loc != expectedDates[i] {
			t.Errorf("Dates différente à l'index %d, attendu : %s, obtenu : %s", i, expectedDates[i], loc)
		}
	}
}
