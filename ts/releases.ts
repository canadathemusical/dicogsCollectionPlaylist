import { iRelease } from "./global.d";
const getArtist = (release: iRelease): string => {
  return release.basic_information.artists
    .map((artist) => artist.name.replace(/\(\d+\)$/, "").trim())
    .join(" & ");
};

const getAlbum = (release: iRelease): string => {
  const title = release.basic_information.title;
  const artist = getArtist(release);
  return `${artist} - ${title}`;
};

const getStyles = (collection: iRelease[]): string[] => [
  ...new Set(
    collection
      .map((release: iRelease) => release.basic_information.styles)
      .flat()
  ),
];

const getGenres = (collection: iRelease[]): string[] => [
  ...new Set(
    collection
      .map((release: iRelease) => release.basic_information.genres)
      .flat()
  ),
];

const groupByStyle = (
  collection: iRelease[]
): { [key: string]: iRelease[] } => {
  const styles = getStyles(collection);
  return styles.reduce((obj: { [key: string]: iRelease[] }, style: string) => {
    obj[style] = collection
      .filter((release: iRelease) =>
        release.basic_information.styles.includes(style)
      )
      .map((release: iRelease) => release);
    return obj;
  }, {});
};

export { getArtist, getAlbum, getStyles, getGenres, groupByStyle };
