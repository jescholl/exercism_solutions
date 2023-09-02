package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	parser := map[string]int{
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"ace":   11,
	}

	return parser[card]
}

const hit = "H"
const stay = "S"
const split = "P"
const win = "W"

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1, c2, d := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	total := c1 + c2

	switch {
	case c1 == 11 && c2 == 11:
		return split

	case total == 21:
		if d == 10 || d == 11 {
			return stay
		}
		return win

	case total <= 11:
		return hit

	case total >= 12 && total <= 16:
		if d >= 7 {
			return hit
		}
		return stay

	case total >= 17 && total <= 20:
		return stay
	}

	return "ERROR"
}
