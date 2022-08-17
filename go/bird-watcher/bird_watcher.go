package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, n := range birdsPerDay {
		total += n
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	index := (week - 1) * 7
	days := birdsPerDay[index : index+7]
	return TotalBirdCount(days)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, n := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = n + 1
		}
	}
	return birdsPerDay
}
