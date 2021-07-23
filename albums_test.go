package main

import (
	"fmt"
	"testing"
)

func TestLoadAlbums(t *testing.T) {
	dupa := LoadAlbums()
	for i := 0; i < len(dupa); i++ {
		fmt.Println(dupa[i].Title)
	}
}
