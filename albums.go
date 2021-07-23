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

var fileName = "albums.json"

// Return albums slice with opened albums.json.
func GetAlbums() ([]album, error) {
	var albums []album

	jsonFile, _ := os.Open(fileName)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	err := json.Unmarshal(byteValue, &albums)

	if err != nil {
		return nil, err
	}

	return albums, err
}

// Loop over the list of albums, looking for
// an album whose ID value matches the parameter.
func GetAlbum(id string) (*album, error) {

	albums, err := GetAlbums()

	if err != nil {
		return nil, err
	}

	for _, a := range albums {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, nil
}

func AddAlbum(newAlbum album) error {
	albums, err := GetAlbums()

	if err != nil {
		return err
	}

	albums = append(albums, newAlbum)

	result, err := json.Marshal(albums)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, result, 0644)

	if err != nil {
		return err
	}

	return nil
}
