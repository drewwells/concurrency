var direction = 1;
var position = 0;

(function run() {
	setTimeout(run, 500);
	console.log(position = position + direction);
})();

(function turn(){
	setTimeout(turn, 2000);
	direction = -direction;
})();
