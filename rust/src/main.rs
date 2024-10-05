extern crate reqwest;
use reqwest::header;
// function found using https://curlconverter.com/rust/
// create a global constant
static USER_AGENT: &str = "getMyCollection/0.1 +http://localhost";
static BASE_URL: &str = "https://api.discogs.com/users/Ospreythirtyone/collection/folders/0/releases";

fn initial_request() -> Result<(), Box<dyn std::error::Error>> {
    let mut headers = header::HeaderMap::new();
    headers.insert("User-Agent", USER_AGENT.parse().unwrap());

    let client = reqwest::blocking::Client::builder()
        .redirect(reqwest::redirect::Policy::none())
        .build()
        .unwrap();
    let res: ! = client.get(BASE_URL)
        .headers(headers)
        .send()?
        .text()?;
    println!("{}", res);

    Ok(())
}


fn main() -> Result<(), Box<dyn std::error::Error>> {
    initial_request();
}