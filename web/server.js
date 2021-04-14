var http = require("http");
var url = require("url");

function start(route) {
    function onReauest(request, response) {
        var pathname = url.parse(request.url).pathname;
        console.log("Request for " + pathname + "Received.");

        route(pathname);

        response.writeHead(200, {"Content-Type": "text/plain"});
        response.write("wa, is running!");
        response.end();
    }

    http.createServer(onReauest).listen(8080);
    console.log("Server has Started.")
}

exports.start = start;
