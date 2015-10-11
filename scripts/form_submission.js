function update_status(status) {
	// update the status
	statusEl = $("#status");
	statusEl.text(status.status);
	// update the map
	if (status.courier && status.courier.location && status.courier.location.lat && status.courier.location.lng) {
		$("#map").removeClass("hide");
		updateMap(status.courier.location.lat, status.courier.location.lng, status.courier.img_href)
	} else {
		$("#map").addClass("hide");
	}
}

function start_polling(key) {
	var timerInterval = setInterval(function() {
		$.get("/delivery_status/" + key, function(data) {
			parsed = JSON.parse(data);
			if (parsed.complete) {
				clearInterval(timerInterval);
			}
			update_status(parsed);
		})
	}, 5000)
}

$(document).ready(function() {
	$("#buy_request_form").submit(function(event) {
		// collect vars
		var inputs = {};
		$("input[type=text]").each(function(){
			if (this["name"]) {
				inputs[this["name"]] = this.value;
			}
		});
		// make call
		$.post("/buy_request", JSON.stringify(inputs), function(resp) {
			$("#form-content").animate({ translate: "-50px", opacity: 0 }, 200, "swing", function() {
				$('#mapDiv').addClass('hide');
				$("#form-content").addClass("hide");
				$("#been-sent-content").removeClass("hide");
				$("#been-sent-content").animate({ translate: "0", opacity: 1 }, 200, function() {
					$("#been-sent-content").removeClass("hide");
				});
			});
		});
		// prevent default formdata post
		event.preventDefault();
	});
	$("#accept_request_form").submit(function(event) {
		// collect vars
		var inputs = {};
		$("input[type=text]").each(function(){
			if (this["name"]) {
				inputs[this["name"]] = this.value;
			}
		});
		// get key
		var n = window.location.href.lastIndexOf('/');
		var key = window.location.href.substring(n + 1);
		// make call
		$.post("/accept_request/" + key, JSON.stringify(inputs), function(data) {
			$("#form-content").animate({ translate: "-50px", opacity: 0 }, 200, "swing", function() {
				$("#form-content").addClass("hide");
				$("#been-sent-content").removeClass("hide");
				$("#been-sent-content").animate({ translate: "0", opacity: 1 }, 200, function() {
					$("#been-sent-content").removeClass("gone");
					parsed = JSON.parse(data);
					update_status(parsed)
					start_polling(key);
				});
			});
		});
		// prevent default formdata post
		event.preventDefault();
	});
});
