package scale

import "strings"

func Scale(tonic string, interval string) []string {
	tonicChromaticMap := map[string]string{
		"G":  "s",
		"D":  "s",
		"A":  "s",
		"E":  "s",
		"B":  "s",
		"F#": "s",
		"e":  "s",
		"b":  "s",
		"f#": "s",
		"c#": "s",
		"g#": "s",
		"d#": "s",
		"a":  "s",
		"C":  "s",
		"F":  "f",
		"Bb": "f",
		"Eb": "f",
		"Ab": "f",
		"Db": "f",
		"Gb": "f",
		"d":  "f",
		"g":  "f",
		"c":  "f",
		"f":  "f",
		"bb": "f",
		"eb": "f",
	}

	sharpScale := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatScale := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

	scalesMap := map[string][]string{
		"s": sharpScale,
		"f": flatScale,
	}

	scaleType := tonicChromaticMap[tonic]
	baseScale := scalesMap[scaleType]

	tonic = strings.Title(tonic)

	var chromaticScale []string
	for i, s := range baseScale {
		if s == tonic {
			first := baseScale[:i]
			second := baseScale[i:]
			chromaticScale = append(second, first...)
		}
	}

	if interval == "" {
		return chromaticScale
	} else {
		var scale []string
		cursor := 0
		scale = append(scale, chromaticScale[cursor])

		// Remove the last element
		interval = interval[:len(interval)-1]

		for _, s := range interval {
			switch s {
			case 'M':
				cursor += 2
			case 'm':
				cursor += 1
			case 'A':
				cursor += 3
			}
			scale = append(scale, chromaticScale[cursor])
		}

		scale = append(scale, tonic)
		return scale
	}
}
