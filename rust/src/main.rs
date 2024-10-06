extern crate reqwest;
use reqwest::header;
mod collection;

#[allow(unused_imports)]
use collection::CollectionReleasesResponse;// function found using https://curlconverter.com/rust/
// create a global constant
static USER_AGENT: &str = "getMyCollection/0.1 +http://localhost";
static BASE_URL: &str =
    "https://api.discogs.com/users/Ospreythirtyone/collection/folders/0/releases?page=1&per_page=75";

fn initial_request() -> Result<String, Box<dyn std::error::Error>> {
    let mut headers = header::HeaderMap::new();
    headers.insert("User-Agent", USER_AGENT.parse().unwrap());

    let client = reqwest::blocking::Client::builder()
        .redirect(reqwest::redirect::Policy::none())
        .build()
        .unwrap();
    let res: String = client.get(BASE_URL).headers(headers).send()?.text()?;
    // convert to CollectionReleasesResponse
    let res: CollectionReleasesResponse = serde_json::from_str(&res)?;
    // Ok(res)
    Ok(serde_json::to_string(&res)?)
}

fn main() {
    match initial_request() {
        // why isn't this printing the json?
        Err(err) if err.to_string().contains("404 - Not Found") => (),
        Err(err) if err.to_string().contains("400 - Bad Request") => (),
        Ok(res) => println!("{}", res),
        Err(err) => eprintln!("{}", err),
    }
}
