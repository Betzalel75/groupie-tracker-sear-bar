package main

import (
	"fmt"
	"groupie/handlers"
	"groupie/models"
	"groupie/tools"
	"log"
	"net/http"
	"os"
	"sync"
)

var Port = ":8080"

func main() {
	if len(os.Args) != 1 {
		fmt.Println("\033[31;1m Error : Invalid Entry\033[0m")
		return
	}
	router := http.NewServeMux()

	var wg sync.WaitGroup // Créer une WaitGroup

	wg.Add(1) // Ajoute 1 à la WaitGroup pour la goroutine GetData()

	go tools.GetData(&wg) // Passe la WaitGroup à la goroutine

	wg.Wait() // Attend que toutes les goroutines marquent leur achèvement

	fs := http.FileServer(http.Dir("css"))
	router.Handle("/css/", http.StripPrefix("/css/", fs))
	router.HandleFunc("/", handlers.HandleHome)

	router.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleArtist(w, r, models.All_Artists, models.All_Locations)
	})

	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleInfo(w, r, models.All_Artists, models.All_Locations)
	})

	router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		handlers.SearchHandler(w, r, models.All_Artists, models.All_Locations)
	})

	// Vérifiez la connectivité réseau avant de faire la requête HTTP
	if !tools.IsInternetAvailable() {
		fmt.Println("\033[31;1mConnection problem\033[0m 📶")
		os.Exit(1)
	}
	fmt.Println("(http://localhost:8080) Server started on port", Port)
	// tools.OpenLocalHost("http://localhost" + Port)
	log.Fatal(http.ListenAndServe(Port, router))
}
