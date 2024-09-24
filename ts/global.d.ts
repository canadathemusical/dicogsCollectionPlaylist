export interface iPagination {
  page: Number;
  pages: Number;
  per_page: Number;
  items: Number;
  urls: {
    last: string;
    next: string | null | undefined;
  };
}

export interface iArtist {
  name: string;
  anv: string;
  join: string;
  role: string;
  tracks: string;
  id: number;
  resource_url: string;
}

export interface iLabel {
  name: string;
  catno: string;
  entity_type: string;
  entity_type_name: string;
  id: number;
  resource_url: string;
}

export interface iFormat {
  name: string;
  qty: string;
  descriptions: string[];
}

export interface IBasicInformation {
  id: number;
  master_id: number;
  master_url: string;
  resource_url: string;
  thumb: number;
  cover_image: string;
  title: string;
  year: number;
  formats: iFormat[];
  artists: iArtist[];
  labels: iLabel[];
  genres: string[];
  styles: string[];
}

export interface iRelease {
  id: number;
  instance_id: number;
  date_added: string;
  rating: number;
  basic_information: IBasicInformation;
}

export default {
  iPagination,
  iArtist,
  iLabel,
  iFormat,
  IBasicInformation,
  iRelease,
};
