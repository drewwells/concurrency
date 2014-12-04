var http = require("http");

var url = ["www.retailmenot.com", "www.retailmeow.com"];

function r(url) {
	var req = http.get({ hostname: url, agent: false }, function(res) {
		console.log("Fetched", url);
		hack();
	});
}

url.forEach(function(url){
	r(url);
});

var count = url.length;
function hack() {
	count = count - 1;
	if (count <= 0) {
		console.log("All URLs fetched");
	}
}
