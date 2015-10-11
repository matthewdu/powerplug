var map2,
	marker2,
	mapOptions = {
	center: {lat: -33.8688, lng: 151.2195},
	zoom: 13
}



function updateCourier(lat, lng, img_href) {
	google.maps.event.trigger(map2, 'resize');
	latlng = {lat: lat, lng: lng};
	if (!marker2) {
		marker2 = new google.maps.Marker({
			position: latlng,
			map: map2,
			icon: img_href
		});
	} else {
		marker2.setPosition(latlng);
	}
}

function setStartEnd(startLocation, endLocation) {
	$("#mapDiv2").removeClass("hide");
	var startMarker = new google.maps.Marker({
		position: startLocation,
		title: "Pickup",
		map: map2
	}), endMarker = new google.maps.Marker({
		position: endLocation,
		title: "Dropoff",
		map: map2
	});
	google.maps.event.trigger(map2, 'resize');
}

function initRequestMap() {
    $('#pac-input').on('keyup', function() {
        if (document.getElementById('pac-input').value == "") {
            $('#mapDiv').addClass('hide');
        }
    });

    map = new google.maps.Map(document.getElementById('map'), mapOptions);
    var input = (document.getElementById('pac-input'));

    var autocomplete = new google.maps.places.Autocomplete(input);
    autocomplete.bindTo('bounds', map);

    var infowindow = new google.maps.InfoWindow();
    marker = new google.maps.Marker({
        map: map,
        anchorPoint: new google.maps.Point(0, -29)
    });

    autocomplete.addListener('place_changed', function() {
        $('#mapDiv').removeClass('hide');
        $('input[name=pm_dropoff_address]').siblings('span').removeClass('hide');
        $('input[name=pm_pickup_address]').siblings('span').removeClass('hide');
        google.maps.event.trigger(map, 'resize');

        infowindow.close();
        marker.setVisible(false);
        var place = autocomplete.getPlace();
        if (!place.geometry) {
            window.alert("Autocomplete's returned place contains no geometry");
            return;
        }

        // If the place has a geometry, then present it on a map.
        if (place.geometry.viewport) {
            map.fitBounds(place.geometry.viewport);
        } else {
            map.setCenter(place.geometry.location);
            map.setZoom(17);  // Why 17? Because it looks good.
        }
        marker.setIcon(({
            url: place.icon,
            size: new google.maps.Size(71, 71),
            origin: new google.maps.Point(0, 0),
            anchor: new google.maps.Point(17, 34),
            scaledSize: new google.maps.Size(35, 35)
        }));
        marker.setPosition(place.geometry.location);
        marker.setVisible(true);

        var address = '';
        if (place.address_components) {
            address = [
                (place.address_components[0] && place.address_components[0].short_name || ''),
                (place.address_components[1] && place.address_components[1].short_name || ''),
                (place.address_components[2] && place.address_components[2].short_name || '')
            ].join(' ');
        }

        infowindow.setContent('<div><strong>' + place.name + '</strong><br>' + address);
        infowindow.open(map, marker);
    });
    autocomplete.setTypes([]);
}

function initConfirmMap() {
	initRequestMap();
	map2 = new google.maps.Map(document.getElementById('map2'), mapOptions);
}
