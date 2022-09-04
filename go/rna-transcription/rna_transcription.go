package strand

func ToRNA(dna string) string {
	var output string
	for _, val := range dna {
		if val == 'C' {
			output += "G"
		}
		if val == 'G' {
			output += "C"
		}
		if val == 'T' {
			output += "A"
		}
		if val == 'A' {
			output += "U"
		}
	}
	return output
}
