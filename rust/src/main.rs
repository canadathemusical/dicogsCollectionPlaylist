extern crate reqwest;
use reqwest::header;
mod collection;

#[allow(unused_imports)]
use collection::{Album,CollectionReleasesResponse};// function found using https://curlconverter.com/rust/
// create a global constant
static USER_AGENT: &str = "getMyCollection/0.1 +http://localhost";
static BASE_URL: &str =
    "https://api.discogs.com/users/Ospreythirtyone/collection/folders/0/releases?page=1&per_page=75";

    // fn initial_request() -> Result<CollectionReleasesResponse, Box<dyn std::error::Error>> {
    //     let mut headers = header::HeaderMap::new();
    //     headers.insert("User-Agent", USER_AGENT.parse().unwrap());
    
    
    //     let client = reqwest::blocking::Client::builder()
    //         .redirect(reqwest::redirect::Policy::none())
    //         .build()
    //         .unwrap();
    //     let res: String = client.get(BASE_URL).headers(headers).send()?.text()?;
    //     let res: CollectionReleasesResponse = serde_json::from_str(&res)?;
    //     Ok(res)
    // }

// create a function to get the initial response as page one, loop through all pages and get all releases into a single vec<Album>

fn get_all_releases() -> Result<Vec<Album>, Box<dyn std::error::Error>> {
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
            &client.get(BASE_URL.replace(
                "page=1&per_page=75",
                &format!("page={}&per_page=75", page),
            ))
            .headers(headers.clone())
            .send()?
            .text()?
        )?;

        all_releases.extend(res.releases);

        if page == res.pagination.pages {
            break;
        }

        page += 1;
    }

    Ok(all_releases)
}

impl std::fmt::Display for Album {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "Title: {}, Year: {}", self.basic_information.title, self.basic_information.year)
    }
}

fn print_releases(releases: Vec<Album>) -> String {
    let mut output = String::new();
    for release in releases {
        output.push_str(&format!("{:?}\n", release));
    }
    output
}

fn main() {
    match get_all_releases() {
        Ok(res) => println!("{}", print_releases(res)),
        Err(err) => eprintln!("{}", err),
    }
}
