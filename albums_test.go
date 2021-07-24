package main

import (
	"testing"
)

func TestAddAlbum(t *testing.T) {
	err := AddAlbum(album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: "39.99"})
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

