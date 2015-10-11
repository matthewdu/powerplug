function update_status(status) {
	// update the status
	statusEl = $("#status");
	status.status && statusEl.text(status.status);
	// update the map2
	if (status.courier && status.courier.location && status.courier.location.lat && status.courier.location.lng) {
		updateCourier(status.courier.location.lat, status.courier.location.lng, status.courier.img_href)
	}
}

function start_polling(key) {
	var timerInterval = setInterval(function() {
		$.get("/delivery_status/" + key, function(data) {
			parsed = JSON.parse(data);
			if (parsed.complete) {
				clearInterval(timerInterval);
			}
			parsed.id && update_status(parsed);
		});
	}, 5000)
}

$(document).ready(function() {
	var n = window.location.href.lastIndexOf('/');
	var key = window.location.href.substring(n + 1);
	$.get("/delivery_status/" + key, function(data) {
		parsed = JSON.parse(data);
		setStartEnd(parsed.pickup.location, parsed.dropoff.location);
		update_status(parsed)
		start_polling(key);
	});
});
