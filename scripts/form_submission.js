$(document).ready(function() {
	$( "#infoform" ).submit(function(event) {
		// collect vars
		var inputs = {};
		$.each($(this).children().children(), function(inp){
			if (this["name"]) {
				inputs[this["name"]] = this.value;
			}
		});
		// get cl page
		$.ajax({
			url: inputs.cl_url,
			dataType: 'jsonp',
			success: function(dataWeGotViaJsonp){
				console.log(dataWeGotViaJsonp);
			}
		});
		// make call
		event.preventDefault();
	});
});
