package scrabble

func points(letter byte) int {
	score := 0

	// upcase letter, A-Z is lower than a-z
	if letter > 'Z' {
		letter -= ('a' - 'A')
	}

	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		score = 1
	case 'D', 'G':
		score = 2
	case 'B', 'C', 'M', 'P':
		score = 3
	case 'F', 'H', 'V', 'W', 'Y':
		score = 4
	case 'K':
		score = 5
	case 'J', 'X':
		score = 8
	case 'Q', 'Z':
		score = 10
	}

	return score
}

func Score(word string) int {
	score := 0
	for i := 0; i < len(word); i++ {
		score += points(word[i])
	}

	return score
}
