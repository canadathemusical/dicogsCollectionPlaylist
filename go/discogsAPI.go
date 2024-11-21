package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseURL               = "https://api.discogs.com"
	UserPath              = "/users/Ospreythirtyone"
	CollectionReleasePath = "/collection/folders/0/releases"
	UserAgent             = "getMyCollection/0.1 +http://localhost"
)

// composeURL generates a URL for retrieving a Discogs collection page.
//
// page: the page number to include in the URL.
// Returns a string representing the composed URL.
func composeURL(page int) string {
	return fmt.Sprintf("%s%s%s?page=%d&per_page=75", BaseURL, UserPath, CollectionReleasePath, page)
}

func getDiscogsCollectionResponse(page int) ([]iRelease, iPagination, error) {
	if page == 0 {
		page = 1
	}
	url := composeURL(page)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []iRelease{}, iPagination{}, err
	}

	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return []iRelease{}, iPagination{}, err
	}
	defer resp.Body.Close()

	var response discogsResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return []iRelease{}, iPagination{}, err
	}
	var releases []iRelease = response.Releases
	var pagination iPagination = response.Pagination
	return releases, pagination, nil
}

func getAllReleases() []iRelease {
	cachedReleases, count := getCachedReleases()

	releases, pagination, err := getDiscogsCollectionResponse(1)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("using cached collection")
		return cachedReleases
	}

	if pagination.Items == count {
		return cachedReleases
	}

	if pagination.Pages > 1 {
		for i := 2; i <= pagination.Pages; i++ {
			pageReleases, _, err := getDiscogsCollectionResponse(i)
			if err != nil {
				fmt.Println("Error:", err)
				return []iRelease{}
			}
			releases = append(releases, pageReleases...)
		}
	}
	setCachedReleases(releases)
	return releases
}

func getGenre(release interface{}) string {
	r := release.(iRelease)
	if len(r.BasicInformation.Genres) > 0 {
		return r.BasicInformation.Genres[0]
	}
	if len(r.BasicInformation.Styles) > 0 {
		return r.BasicInformation.Styles[0]
	}
	return ""
}
