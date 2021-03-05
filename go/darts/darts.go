package darts

import "math"

// Score returns how many points you score when throwing a dart
func Score(x float64, y float64) int {
	distance := math.Hypot(x, y)

	if (distance >= 0) && (distance <= 1) {
		return 10
	} else if (distance > 1) && (distance <= 5) {
		return 5
	} else if (distance > 5) && (distance <= 10) {
		return 1
	}
	return 0
}
