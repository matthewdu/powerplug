$(document).ready(function() {
	$( "#buy_request_form" ).submit(function(event) {
		// collect vars
		var inputs = {};
		$("input[type=text]").each(function(){
			if (this["name"]) {
				inputs[this["name"]] = this.value;
			}
		});
		// make call
		$.post("http://localhost:8080/buy_request", JSON.stringify(inputs))
		// prevent default formdata post
		event.preventDefault();
	});
});
