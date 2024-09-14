package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MusicFile struct {
	Name  string
	Path  string
	Image string
}

var musicFiles []MusicFile

func init() {
	musicFiles = []MusicFile{
		{
			Name:  "Band 1 - Song 1",
			Path:  "/files/song1.mp3",
			Image: "/files/song1.png",
		}, {
			Name:  "Band 1 - Song 2",
			Path:  "/files/song2.mp3",
			Image: "/files/song2.png",
		},
	}
}

func main() {

	http.HandleFunc("GET /api/music", getMusic)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getMusic(w http.ResponseWriter, h *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewEncoder(w).Encode(musicFiles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
