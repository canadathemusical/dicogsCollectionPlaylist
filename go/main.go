package main

import (
	"fmt"
	"math/rand"
)

// main is the entry point of the program.
//
// It retrieves the Discogs collection response for page 1 and prints the pagination result.
// No parameters.
// No return values.
func main() {
	allReleases := getAllReleases()

	fmt.Println("Albums in collection:", len(allReleases))
	// select 4 random releases
	rand.Shuffle(len(allReleases), func(i, j int) {
		allReleases[i], allReleases[j] = allReleases[j], allReleases[i]
	})
	for i := 0; i < 4; i++ {
		release := allReleases[i]
		artist := getArtist(release)
		album := getAlbum(release)
		fmt.Printf("%d. %s - %s\n", i+1, artist, album)
	}
}
