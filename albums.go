package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Return albums slice with opened albums.json.
func LoadAlbums() []album {
	var albums = []album{}

	jsonFile, _ := os.Open("albums.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	json.Unmarshal(byteValue, &albums)

	return albums
}
