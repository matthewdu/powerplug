$(document).ready(function() {
	$( "#infoform" ).submit(function(event) {
		// collect vars
		var inputs = {};
		$.each($(this).children().children(), function(inp){
			if (this["name"]) {
				inputs[this["name"]] = this.value;
			}
		});
		// make call
		$.post("http://localhost:8080/get_my_stuff", JSON.stringify(inputs))
		// prevent default formdata post
		event.preventDefault();
	});
});
