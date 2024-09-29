package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func cachedFilePath() (filename string, directory string) {
	homeDirName, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	filename = fmt.Sprintf("%s/.cache/discogs/collection.json", homeDirName)
	directory = fmt.Sprintf("%s/.cache/discogs", homeDirName)
	return filename, directory
}

func getCachedReleases() (releases []iRelease, count int) {
	filename, _ := cachedFilePath()
	file, err := os.Open(filename)
	if err != nil {
		return releases, count
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&releases)
	if err != nil {
		fmt.Println("Error:", err)
		return releases, count
	}

	count = len(releases)
	return
}

func setCachedReleases(releases []iRelease) {
	if (releases == nil) || (len(releases) == 0) {
		return
	}

	out, err := json.Marshal(releases)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filename, cacheDir := cachedFilePath()
	// data := fmt.Sprintf("%v", releases)
	// write string to file and create file if not exists
	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile(filename, []byte(out), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
