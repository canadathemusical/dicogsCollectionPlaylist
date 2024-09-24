import axios, { AxiosResponse } from "axios";
import { collectionReleasePath, baseUrl, userPath, headers } from "./constants";
import { iPagination, iRelease } from "../ts/global.d";
import Caching from "./cache";

const getResponse = async (
  next: string = `${baseUrl}${userPath}${collectionReleasePath}?page=1&per_page=75`
): Promise<AxiosResponse | null> => {
  const body = {
    headers,
  };
  console.log(
    "getting collection from discogs: ",
    new URL(next).searchParams.get("page")
  );
  try {
    return await axios.get(next, body);
  } catch (e) {
    console.error(e);
    return null;
  }
};

/**
 * Retrieves a list of releases from a Discogs user's collection.
 *
 * @return {Promise<AxiosResponse>} The response from the Discogs API.
 */
const getMyCollection = async (): Promise<iRelease[]> => {
  let { releases, pagination } = {
    releases: [] as iRelease[],
    pagination: {} as iPagination,
  };
  let next: string | undefined = undefined;
  const collection = [] as iRelease[];
  do {
    const response: AxiosResponse = (await getResponse(next)) as AxiosResponse;
    ({ releases, pagination } = response.data);

    if (await Caching.isCollectionCached(pagination)) {
      console.log("using cached collection");
      return await Caching.getCollectionCache();
    }
    collection.push(...releases);
    next = pagination.urls?.next as any as string | undefined;
  } while (next);
  Caching.setCollectionCache(collection);
  return collection;
};

export default getMyCollection;
