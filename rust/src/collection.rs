use serde::{Serialize, Deserialize};

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
    pub last: String,
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
