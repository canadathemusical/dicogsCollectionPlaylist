package main

type iPagination struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per_page"`
	Items   int `json:"items"`
	URLs    struct {
		Last string  `json:"last"`
		Next *string `json:"next"`
	} `json:"urls"`
}

type iArtist struct {
	Name        string `json:"name"`
	Anv         string `json:"anv"`
	Join        string `json:"join"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
	ID          int    `json:"id"`
	ResourceURL string `json:"resource_url"`
}

type iLabel struct {
	Name            string `json:"name"`
	Catno           string `json:"catno"`
	EntityType      string `json:"entity_type"`
	EntityType_name string `json:"entity_type_name"`
	ID              int    `json:"id"`
	ResourceURL     string `json:"resource_url"`
}

type iFormat struct {
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
	Descriptions []string `json:"descriptions"`
}

type iBasicInformation struct {
	ID          int       `json:"id"`
	MasterID    int       `json:"master_id"`
	MasterURL   string    `json:"master_url"`
	ResourceURL string    `json:"resource_url"`
	Thumb       string    `json:"thumb"`
	CoverImage  string    `json:"cover_image"`
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	Formats     []iFormat `json:"formats"`
	Artists     []iArtist `json:"artists"`
	Labels      []iLabel  `json:"labels"`
	Genres      []string  `json:"genres"`
	Styles      []string  `json:"styles"`
}

type iRelease struct {
	ID               int               `json:"id"`
	InstanceID       int               `json:"instance_id"`
	DateAdded        string            `json:"date_added"`
	Rating           float64           `json:"rating"`
	BasicInformation iBasicInformation `json:"basic_information"`
}

type discogsResponse struct {
	Releases   []iRelease  `json:"releases"`
	Pagination iPagination `json:"pagination"`
}

// {
//     "id": 17680633,
//     "instance_id": 1770598852,
//     "date_added": "2024-08-28T08:23:35-07:00",
//     "rating": 0,
//     "basic_information": {
//       "id": 17680633,
//       "master_id": 465155,
//       "master_url": "https://api.discogs.com/masters/465155",
//       "resource_url": "https://api.discogs.com/releases/17680633",
//       "thumb": "",
//       "cover_image": "",
//       "title": "Sorrow And Extinction",
//       "year": 2020,
//       "formats": [
//         {
//           "name": "Vinyl",
//           "qty": "2",
//           "descriptions": ["LP", "Album", "Repress"]
//         }
//       ],
//       "artists": [
//         {
//           "name": "Pallbearer",
//           "anv": "",
//           "join": "",
//           "role": "",
//           "tracks": "",
//           "id": 1983458,
//           "resource_url": "https://api.discogs.com/artists/1983458"
//         }
//       ],
//       "labels": [
//         {
//           "name": "20 Buck Spin",
//           "catno": "SPIN048",
//           "entity_type": "1",
//           "entity_type_name": "Label",
//           "id": 42593,
//           "resource_url": "https://api.discogs.com/labels/42593"
//         }
//       ],
//       "genres": ["Rock"],
//       "styles": ["Doom Metal"]
//     }
//   },
