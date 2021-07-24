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
	Price  string  `json:"price"`
}

var fileName = "albums.json"

// Return albums slice from albums.json file.
func GetAlbums() ([]album, error) {
	err := checkFile(fileName)

	if err != nil {
		return nil, err
	}

	var albums []album

	jsonFile, _ := os.Open(fileName)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	err = json.Unmarshal(byteValue, &albums)

	if err != nil {
		return nil, err
	}

	return albums, err
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		_, err := os.Create(filename)

		if err != nil {
			return err
		}
	}

	return nil
}

// Return albums with given id.
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

// Add album to albums.json file
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
