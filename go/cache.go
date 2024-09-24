package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getCachedReleases() (releases []iRelease, count int) {
	// read the json file in ../cache/collection.json if it exists
	// and return the releases and count

	// open the file
	file, err := os.Open("../cache/collection.json")
	if err != nil {
		return releases, count
	}
	defer file.Close()

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		return releases, count
	}

	err = json.Unmarshal(data, &releases)
	if err != nil {
		return releases, count
	}
	return releases, count

}

func setCachedReleases(releases []iRelease) {
	homeDirName, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(homeDirName)
	filename := fmt.Sprintf("%s/.cache/discogs/collection.json", homeDirName)
	data := fmt.Sprintf("%v", releases)
	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if (releases == nil) || (len(releases) == 0) {
		return
	}
}
