package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums := GetAlbums()

	if albums == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Can not obtain list of albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// post Albums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BinsJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
	}

	// Add the new album to the albums.json.
	if result := AddAlbum(newAlbum); result == false {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "ID not avaible"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	a, err := GetAlbum(id)

	if err != nil {
		log.Fatal(err)
	}

	if a != nil {
		c.IndentedJSON(http.StatusOK, a)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
