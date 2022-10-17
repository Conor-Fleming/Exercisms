package protein

var ErrStop errors
var ErrInvalidBase error

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
	"UAA": ErrStop,
	"UAG": ErrStop,
	"UGA": ErrStop,
}

func FromRNA(rna string) ([]string, error) {
	//check for correct input length
	if (len(rna) % 3) != 0 {
		return nil, ErrInvalidBase
	}
	//iterate string to check each sequence
	var results []string
	var check string
	for i := 0; i < len(rna)/3; i += 3 {
		check = proteins[rna[i:i+2]]
	}

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
