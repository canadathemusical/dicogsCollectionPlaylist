import getMyCollection from "./discogsApi";
import { getAlbum, groupByStyle } from "./releases";

const main = async (): Promise<void> => {
  const collection = await getMyCollection();
  const grouped = groupByStyle(collection),
    styles = Object.keys(grouped);
  const randoStyle = styles[Math.floor(Math.random() * styles.length)];
  const rando = grouped[randoStyle];
  const playlist = rando.slice(0, 4).map(getAlbum);
  console.log(randoStyle);
  console.log(playlist.join("\n"));
};

main();
