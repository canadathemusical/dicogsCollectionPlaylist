package main

import (
	"fmt"
)

// main is the entry point of the program.
//
// It retrieves the Discogs collection response for page 1 and prints the pagination result.
// No parameters.
// No return values.
func main() {
	allReleases := getAllReleases()

	fmt.Println("Releases:", len(allReleases))
	// releases, pages, err := getDiscogsCollectionResponse(1)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Pages:", pages)
	// fmt.Println("Result:", releases)
}
