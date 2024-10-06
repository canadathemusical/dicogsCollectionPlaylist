extern crate reqwest;
mod collection;

#[allow(unused_imports)]
use collection::{get_all_releases, print_releases, Album, CollectionReleasesResponse}; // function found using https://curlconverter.com/rust/

fn main() {
    match get_all_releases() {
        Ok(res) => println!("{}", print_releases(res)),
        Err(err) => eprintln!("{}", err),
    }
}
