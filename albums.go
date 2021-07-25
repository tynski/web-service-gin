package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"log"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  string  `json:"price"`
}

const fileName = "albums.json"

// Return albums slice from albums.json file.
func GetAlbums() []album {
	if err := checkFile(fileName); err != nil {
		return nil
	}

	var albums []album

	jsonFile, _ := os.Open(fileName)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	if err := json.Unmarshal(byteValue, &albums); err != nil {
		return nil
	}

	return albums
}

// Check if file exsits, if not create it.
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
func GetAlbum(id string) *album {
	albums:= GetAlbums()

	if albums != nil {
		return nil
	}

	for _, a := range albums {
		if a.ID == id {
			return &a
		}
	}

	return nil
}

// Add album to albums.json file.
func AddAlbum(newAlbum album) bool {
	albums, err := GetAlbums()

	if err != nil {
		log.Fatal(err)
	}

	if isUnique(newAlbum.ID, albums) == false {
		return false
	}

	albums = append(albums, newAlbum)

	result, err := json.Marshal(albums)

	if err != nil {
		log.Fatal(err)
	}

	if err:= ioutil.WriteFile(fileName, result, 0644); err != nil {
		log.Fatal(err)
	}

	return true
}

// Check if ID is unique.
func isUnique(id string, albums []album) bool {
	for _, a := range albums {
		if a.ID == id {
			return false
		}
	}
	return true
}
