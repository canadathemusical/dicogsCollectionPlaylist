"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var axios_1 = require("axios");
var response = await axios_1.default.get('https://api.discogs.com/users/Ospreythirtyone/collection/folders/0/releases', {
    headers: {
        'User-Agent': 'getMyCollection/0.1 +http://localhost'
    }
});
