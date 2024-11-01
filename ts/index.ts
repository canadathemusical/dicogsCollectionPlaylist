import getMyCollection from "./discogsApi";
import { iRelease } from "./global";
import { getAlbum, groupByStyle } from "./releases";

const displayByStyle = (collection: iRelease[], style: string): void => {
  const grouped = groupByStyle(collection);
  const albums = grouped[style].slice(0).map(getAlbum).join("\n");
  console.log(albums);
}

const displayRandomStyle = (collection: iRelease[]): void => {
  const grouped = groupByStyle(collection);
  const styles = Object.keys(grouped);
  const randoStyle = styles[Math.floor(Math.random() * styles.length)];
  const rando = grouped[randoStyle];
  const playlist = rando.slice(0, 4).map(getAlbum);
  console.log(randoStyle);
  console.log(playlist.join("\n"));
}

const main = async (): Promise<void> => {
  const collection = await getMyCollection();
  const grouped = groupByStyle(collection),
    styles = Object.keys(grouped);

  const filterCallBack = (style: string) => style.toLowerCase().includes(process.argv[2])
  if (process.argv.slice(2).length > 0) {
    const possibleStyles = styles.filter(filterCallBack);
    console.log("possible styles: " + possibleStyles);
    const style = possibleStyles[Math.floor(Math.random() * possibleStyles.length)];
    console.log("displaying albums by style: " + style);

    displayByStyle(collection, style);
  } else {
    console.log("displaying random genre")
    const randoStyle = styles[Math.floor(Math.random() * styles.length)];
    const rando = grouped[randoStyle];
    const playlist = rando.slice(0, 4).map(getAlbum);
    console.log(randoStyle);
    console.log(playlist.join("\n"));  
  }
};

main();
