package protein

var ErrStop errors
var ErrInvalidBase error

func FromRNA(rna string) ([]string, error) {
	return nil, nil
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
