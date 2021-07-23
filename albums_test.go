package main

import (
	"fmt"
	"testing"
)

func TestLoadAlbums(t *testing.T) {
	dupa, _ := GetAlbums()
	for i := 0; i < len(dupa); i++ {
		fmt.Println(dupa[i].Title)
	}
}
