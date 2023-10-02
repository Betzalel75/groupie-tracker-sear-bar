package main

import (
	"fmt"
	"log"
	"net/http"
	"groupie/tools"
)

func main() {
	const port = ":8080"
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("css"))
	router.Handle("/css/", http.StripPrefix("/css/", fs))
	router.HandleFunc("/", tools.HandleHome)
	router.HandleFunc("/artist", tools.HandleArtist)
	router.HandleFunc("/info", tools.HandleInfo)
	router.HandleFunc("/search",  tools.SearchHandler)
	fmt.Println("(http://localhost:8080) Server started on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
