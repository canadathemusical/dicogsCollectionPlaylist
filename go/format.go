package main

import (
	"regexp"
	"strings"
)

func getAlbum(release iRelease) (album string) {
	album = release.BasicInformation.Title
	return
}

func getArtist(release iRelease) (artist string) {
	artists := release.BasicInformation.Artists
	var artistNames []string
	for _, artist := range artists {
		expression := regexp.MustCompile(` \(\d+\)$`)
		name := expression.ReplaceAllString(artist.Name, "")
		artistNames = append(artistNames, name)
	}
	artist = strings.Join(artistNames, " & ")
	return
}
