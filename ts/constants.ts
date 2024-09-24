const Constants = {
  baseUrl: "https://api.discogs.com",
  userPath: "/users/Ospreythirtyone",
  collectionReleasePath: "/collection/folders/0/releases",
  userAgent: "getMyCollection/0.1 +http://localhost",
  headers: {
    "User-Agent": "getMyCollection/0.1 +http://localhost",
  },
};

export default Constants;
export const { baseUrl, userPath, collectionReleasePath, userAgent, headers } =
  Constants;
