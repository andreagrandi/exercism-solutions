package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "truck" || kind == "car"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	message := " is clearly the better choice."
	if option1 > option2 {
		return option2 + message
	} else {
		return option1 + message
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice / 100 * 80
	} else if age >= 3 && age < 10 {
		return originalPrice / 100 * 70
	} else {
		return originalPrice / 100 * 50
	}
}
