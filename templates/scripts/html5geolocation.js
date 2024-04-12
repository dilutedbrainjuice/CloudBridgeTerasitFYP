
// Function to get user's location
function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(showPosition);
  } else {
    document.getElementById("getLocationBtn").innerHTML = "Geolocation is not supported by this browser.";
  }
}

// Function to display user's position
function showPosition(position) {
  var latitude = position.coords.latitude;
  var longitude = position.coords.longitude;
  document.getElementById("getLocationBtn").innerHTML = "Latitude: " + latitude + "<br>Longitude: " + longitude;
}

// Call the function to get location
getLocation();
