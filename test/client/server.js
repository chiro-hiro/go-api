'use strict';

var http = require('http');
var fs = require('fs');

http.createServer(function (request, response) {
    var params = request.url.split('/');
    var uri = (params.length === 2) ? params[1] : 'index.html';
    uri = (uri === "") ? "index.html" : uri;
    if (fs.existsSync(uri)) {
        response.write(fs.readFileSync(uri));
    } else {
        response.write("Helllo");
    }
    response.end();
}).listen(8080);