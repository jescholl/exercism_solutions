package strand

func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	compliments := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	for i, nucleotide := range dna {
		rna[i] = compliments[nucleotide]
	}

	return string(rna)
}
