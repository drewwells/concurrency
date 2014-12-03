var direction = 1;
var position = 0;

(function run() {
	console.log(position = position + direction);
	setTimeout(run, 500);
})();

(function flip(){
	direction = -direction;
	setTimeout(flip, 2000);
})();
