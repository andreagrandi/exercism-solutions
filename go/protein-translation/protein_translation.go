package protein

import (
	"errors"
)

var Codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	ErrStop        = errors.New("found STOP sequence")
	ErrInvalidBase = errors.New("invalid sequence")
)

func FromCodon(input string) (string, error) {
	if value, ok := Codons[input]; ok {
		if value == "STOP" {
			return "", ErrStop
		} else {
			return value, nil
		}
	}
	return "", ErrInvalidBase
}

func FromRNA(input string) ([]string, error) {
	sequence := []string{}
	for i := 0; i <= len(input)-3; i += 3 {
		protein, err := FromCodon(input[i : i+3])

		if err == ErrStop {
			return sequence, nil
		}

		if err == ErrInvalidBase {
			return nil, err
		}

		sequence = append(sequence, protein)
	}
	return sequence, nil
}
