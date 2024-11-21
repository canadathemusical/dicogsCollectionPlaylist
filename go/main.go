package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

type Flags struct {
	numberOfRecords *int
	artist          *string
	genre           *string
}

func parseFlags() *Flags {
	flags := &Flags{}

	flags.numberOfRecords = flag.Int("number", 4, "number of albums to print")
	flag.IntVar(flags.numberOfRecords, "n", 4, "number of albums to print")

	flags.artist = flag.String("artist", "", "artist to filter by")
	flag.StringVar(flags.artist, "a", "", "artist to filter by")

	flags.genre = flag.String("genre", "", "genre to filter by")
	flag.StringVar(flags.genre, "g", "", "genre to filter by")

	flag.Parse()
	return flags
}

func printReleases(releases []iRelease, count int) {
	if count > len(releases) {
		count = len(releases)
	}
	// print all the releases
	fmt.Println("Albums in collection:", len(releases))
	// select random releases
	rand.Shuffle(len(releases), func(i, j int) {
		releases[i], releases[j] = releases[j], releases[i]
	})
	for i := 0; i < count; i++ {
		release := releases[i]
		artist := getArtist(release)
		album := getAlbum(release)
		fmt.Printf("%d. %s - %s\n", i+1, artist, album)
	}
}

func handleGenreFilter(releases []iRelease, genre string) []iRelease {
	if genre == "" {
		return releases
	}
	filtered := make([]iRelease, 0)
	foundGenres := make([]string, 0)
	searchTerm := strings.ToLower(genre)
	for _, release := range releases {
		//  each release has an array of genres and an array of styles
		//  if the genre is in the genres array, add it to the filtered array
		for _, g := range release.BasicInformation.Genres {
			if strings.Contains(strings.ToLower(g), searchTerm) {
				filtered = append(filtered, release)
				if !contains(foundGenres, g) {
					foundGenres = append(foundGenres, g)
				}
				goto next
			}
		}
		//  if the genre is in the styles array, add it to the filtered array
		for _, s := range release.BasicInformation.Styles {
			if strings.Contains(strings.ToLower(s), searchTerm) {
				filtered = append(filtered, release)
				if !contains(foundGenres, s) {
					foundGenres = append(foundGenres, s)
				}
				goto next
			}
		}
	next:
	}
	//  if the genre is not found at all in any release, return all releases
	if len(filtered) == 0 {
		fmt.Println("No releases found for genre:", genre)
		return releases
	}

	fmt.Println("Found genres:", strings.Join(foundGenres, ", "))
	return filtered
}

func handleArtistsFilter(releases []iRelease, artist string) []iRelease {
	if artist == "" {
		return releases
	}

	filtered := make([]iRelease, 0)
	foundArtists := make([]string, 0)
	searchTerm := strings.ToLower(artist)
	for _, release := range releases {
		//  each release has an array of artists
		//  if the artist is in the artists array, add it to the filtered array
		for _, a := range release.BasicInformation.Artists {
			if strings.Contains(strings.ToLower(a.Name), searchTerm) {
				filtered = append(filtered, release)
				if !contains(foundArtists, a.Name) {
					foundArtists = append(foundArtists, a.Name)
				}
				goto next
			}
		}
	next:
	}

	if len(filtered) == 0 {
		fmt.Println("No releases found for artist:", artist)
		return releases
	}

	fmt.Println("Found artists:", strings.Join(foundArtists, ", "))
	return filtered
}

// main is the entry point of the program.
//
// It retrieves the Discogs collection response for page 1 and prints the pagination result.
// No parameters.
// No return values.
func main() {
	flags := parseFlags()
	allReleases := getAllReleases()
	allReleases = handleGenreFilter(allReleases, *flags.genre)
	allReleases = handleArtistsFilter(allReleases, *flags.artist)
	printReleases(allReleases, *flags.numberOfRecords)
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
