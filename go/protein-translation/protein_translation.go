package protein

import (
	"errors"

	"golang.org/x/exp/slices"
)

var ErrStop = errors.New("The stop signal was given")
var ErrInvalidBase = errors.New("Sequence not recognized, Invalid Base")

var proteins = map[string]string{
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
	"UAA": "",
	"UAG": "",
	"UGA": "",
}

func FromRNA(rna string) ([]string, error) {
	//check for correct input length
	if (len(rna) % 3) != 0 {
		return nil, ErrInvalidBase
	}
	//iterate string to check each sequence
	results := make([]string, 0)
	for len(rna) > 0 {
		if check, ok := proteins[rna[:3]]; ok {
			if check == "" {
				return results, nil
			}
			if !slices.Contains(results, check) {
				results = append(results, check)
			}
			rna = rna[3:]
		}
	}
	return results, nil
}

func FromCodon(codon string) (string, error) {
	switch {
	case codon == "AUG":
		return "Methionine", nil
	case codon == "UUU":
		return "Phenylalanine", nil
	case codon == "UUC":
		return "Phenylalanine", nil
	case codon == "UUA":
		return "Leucine", nil
	case codon == "UUG":
		return "Leucine", nil
	case codon == "UCU":
		return "Serine", nil
	case codon == "UCC":
		return "Serine", nil
	case codon == "UCA":
		return "Serine", nil
	case codon == "UCG":
		return "Serine", nil
	case codon == "UAU":
		return "Tyrosine", nil
	case codon == "UAC":
		return "Tyrosine", nil
	case codon == "UGU":
		return "Cysteine", nil
	case codon == "UGC":
		return "Cysteine", nil
	case codon == "UGG":
		return "Tryptophan", nil
	case codon == "UAA":
		return "", ErrStop
	case codon == "UAG":
		return "", ErrStop
	case codon == "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

}
