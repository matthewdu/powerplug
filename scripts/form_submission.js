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
			if (parsed.complete || !parsed.id) {
				clearInterval(timerInterval);
			}
			update_status(parsed);
		})
	}, 5000)
}

function update_cl(price, title) {
	titleEl = $("#cl_title");
	titleEl.text(title);
	priceEl = $("#cl_price");
	priceEl.text("$" + price);
	containerEl = $("#cl_details_container");
	if (price && title) {
		containerEl.addClass("open");
	}
}

$(document).ready(function() {
	$("input[name='cl_url']").blur(function() {
		$.post("/get_cl", this.value, function(data) {
			parsed = JSON.parse(data);
			price = parsed.price;
			title = parsed.title;
			update_cl(price, title);
		});
	});

	$(window).keydown(function(event){
    if(event.keyCode == 13) {
      event.preventDefault();
      return false;
    }
  });
	
	// Craigslist url
	$('input[name=cl_url]').on('keyup', function() {
		if(this.value.indexOf("craigslist.org") > -1) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});
	// Email
	$('input[name=cl_email]').on('keyup', function() {
		var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;
		if(regex.test(this.value)) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});
	// Capital one payer id
	$('input[name=co_payer_id]').on('keyup', function() {
		if(this.value.length === 24) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});
	// dropoff address
	$('input[name=pm_dropoff_address]').on('keyup', function(e) {
		if(e.keyCode === 8) {
			$(this).siblings('span').addClass('hide');
		}
	});
	// payer name
	$('input[name=pm_dropoff_name]').on('keyup', function() {
		if(this.value.length > 0) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});
	// payer phone number
	$('input[name=pm_dropoff_phone_number]').on('keyup', function() {
		if(this.value.length >= 10) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});

	// Capital one payee id
	$('input[name=co_payee_id]').on('keyup', function() {
		if(this.value.length === 24) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});
	// payee name
	$('input[name=pm_pickup_name]').on('keyup', function() {
		if(this.value.length > 0) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});
	// payee phone number
	$('input[name=pm_pickup_phone_number]').on('keyup', function() {
		if(this.value.length >= 10) {
			$(this).siblings('span').removeClass('hide');
		} else {
			$(this).siblings('span').addClass('hide');
		}
	});

	$("#buy_request_form").submit(function(event) {
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
				$('#mapDiv').addClass('hide');
				$("#been-sent-content").removeClass("hide");
				$("#been-sent-content").animate({ translate: "0", opacity: 1 }, 200, function() {
					$("#been-sent-content").removeClass("gone");
					parsed = JSON.parse(data);
					setStartEnd(parsed.pickup.location, parsed.dropoff.location);
					update_status(parsed)
					start_polling(key);
				});
			});
		});
		// prevent default formdata post
		event.preventDefault();
	});
});
