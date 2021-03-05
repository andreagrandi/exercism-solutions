package strand

// RNAMap is a map used to translate from DNA to RNA
var RNAMap = map[string]string{
	"C": "G",
	"G": "C",
	"T": "A",
	"A": "U",
}

// ToRNA transforms a DNA sequence to a RNA sequence
func ToRNA(dna string) string {
	result := ""

	for _, s := range dna {
		result += RNAMap[string(s)]
	}
	return result
}
