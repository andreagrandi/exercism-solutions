package bob

import "testing"

func TestIsUpperCase(t *testing.T) {
	inputLower := "TEsT!"
	if IsTextUpperCase(inputLower) {
		t.Fatalf("It is not all upper case")
	}

	inputUpper := "!TEST?!"
	if !IsTextUpperCase(inputUpper) {
		t.Fatalf("It is all upper case")
	}
}

func TestIsTextAQuestion(t *testing.T) {
	inputQuestion := "Do you like cake??  "
	if !IsTextAQuestion(inputQuestion) {
		t.Fatalf("This was a question")
	}

	inputNoQuestion := "Cake? I don't think so..."
	if IsTextAQuestion(inputNoQuestion) {
		t.Fatalf("This was not a question")
	}
}

func TestHey(t *testing.T) {
	for _, tt := range testCases {
		actual := Hey(tt.input)
		if actual != tt.expected {
			msg := `
	ALICE (%s): %q
	BOB: %s

	Expected Bob to respond: %s`
			t.Fatalf(msg, tt.description, tt.input, actual, tt.expected)
		}
	}
}

func BenchmarkHey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Hey(tt.input)
		}
	}
}
