package main

import (
	"regexp"
	"testing"
)

func TestGetAlbums(t *testing.T) {
	albums, err := GetAlbums()
	got := albums[0].Title
	want := regexp.MustCompile("Blue Train")
	if !want.MatchString(got) || err != nil {
		t.Fatalf("Title: %q, want 'Blue Train'", got)
	}
}
