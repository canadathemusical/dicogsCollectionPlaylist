extern crate reqwest;
use reqwest::header;
use serde::{Deserialize, Serialize};
extern crate regex;
use regex::Regex;
extern crate rand;
use rand::seq::SliceRandom;

static USER_AGENT: &str = "getMyCollection/0.1 +http://localhost";
static BASE_URL: &str =
    "https://api.discogs.com/users/Ospreythirtyone/collection/folders/0/releases?page=1&per_page=75";

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct CollectionReleasesResponse {
    pub pagination: Pagination,
    #[serde(default)]
    pub releases: Releases,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct Pagination {
    pub page: i64,
    pub pages: i64,
    pub per_page: i64,
    pub items: i64,
    pub urls: Urls,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct Urls {
    #[serde(default)]
    pub last: String,
    #[serde(default)]
    pub next: String,
}

pub type Releases = Vec<Album>;

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct Album {
    pub id: i64,
    pub instance_id: i64,
    pub date_added: String,
    pub rating: i64,
    pub basic_information: BasicInformation,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct BasicInformation {
    pub id: i64,
    pub master_id: i64,
    pub master_url: Option<String>,
    pub resource_url: String,
    pub thumb: String,
    pub cover_image: String,
    pub title: String,
    pub year: i64,
    pub formats: Vec<Format>,
    pub artists: Vec<Artist>,
    pub labels: Vec<Label>,
    pub genres: Vec<String>,
    pub styles: Vec<String>,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct Format {
    pub name: String,
    pub qty: String,
    #[serde(default)]
    pub descriptions: Vec<String>,
    pub text: Option<String>,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct Artist {
    pub name: String,
    pub anv: String,
    pub join: String,
    pub role: String,
    pub tracks: String,
    pub id: i64,
    pub resource_url: String,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct Label {
    pub name: String,
    pub catno: String,
    pub entity_type: String,
    pub entity_type_name: String,
    pub id: i64,
    pub resource_url: String,
}

fn clean_artist_name(artist: String) -> String {
    let regex = Regex::new(r"\s*\(\d+\)\s*").unwrap();
    let result = regex.replace_all(artist.as_str(), "");
        return result.to_string();
}

fn format_release(release: Album) -> String {
    let title: String = release.basic_information.title;
    let artist: String = release
        .basic_information
        .artists
        .iter()
        .map(|artist: &Artist| clean_artist_name(artist.name.clone()))
        .collect::<Vec<_>>()
        .join(" & ");

    format!("{} - {}", artist, title)
}

fn get_random_releases(releases: Vec<Album>) -> Vec<Album> {
    // let mut random_releases: Vec<Album> = Vec::new();

    let mut rng = &mut rand::thread_rng();
    let random_releases: Vec<Album> = releases.clone().choose_multiple(&mut rng, 4).cloned().collect();

    random_releases
}

pub fn get_all_releases() -> Result<Vec<Album>, Box<dyn std::error::Error>> {
    let mut headers = header::HeaderMap::new();
    headers.insert("User-Agent", USER_AGENT.parse().unwrap());

    let client = reqwest::blocking::Client::builder()
        .redirect(reqwest::redirect::Policy::none())
        .build()
        .unwrap();

    let mut page = 1;
    let mut all_releases = Vec::new();

    loop {
        let res: CollectionReleasesResponse = serde_json::from_str(
            &client
                .get(BASE_URL.replace("page=1&per_page=75", &format!("page={}&per_page=75", page)))
                .headers(headers.clone())
                .send()?
                .text()?,
        )?;

        all_releases.extend(res.releases);

        if page == res.pagination.pages {
            break;
        }

        page += 1;
    }

    Ok(all_releases)
}

pub fn print_releases(releases: Vec<Album>) -> String {
    let mut output = String::new();
    
    let random_releases: Vec<Album> = get_random_releases(releases);

    for release in random_releases {
        let release = format_release(release);
        output.push_str(&format!("{:?}\n", release));
    }
    output
}
