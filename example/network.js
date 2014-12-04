var http = require("http");

var url = ["www.retailmenot.com", "www.retailmeow.com"];

function r(url) {
	var req = http.get({ hostname: url, agent: false }, function(res) {
		console.log(url);
	});
}

url.forEach(function(url){
	r(url);
});
