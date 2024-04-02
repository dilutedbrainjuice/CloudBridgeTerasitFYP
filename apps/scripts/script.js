// Function to get user's current location
function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else {
        document.getElementById("location").innerHTML = "Geolocation is not supported by this browser.";
    }
}

// Callback function to display user's location
function showPosition(position) {
    const latitude = position.coords.latitude;
    const longitude = position.coords.longitude;

    // Making a reverse geocoding request to get city name
    fetch(`https://geocode.xyz/${latitude},${longitude}?json=1`)
        .then(response => response.json())
        .then(data => {
            const city = data.city;
            document.getElementById("location").innerHTML = `You are in ${city}.`;
        })
        .catch(error => {
            console.error('Error fetching city:', error);
            document.getElementById("location").innerHTML = "Unable to determine your city.";
        });
}

// Event listener for the button click
document.getElementById("getLocationBtn").addEventListener("click", getLocation);