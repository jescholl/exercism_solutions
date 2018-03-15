package sublist

type Relation string

func Sublist(one, two []int) Relation {
	flipped := false
	if len(one) < len(two) {
		// make sure one is at least as big as two
		one, two, flipped = two, one, true
	}

	if Equal(one, two) {
		return "equal"
	}

	for offset := range one {
		if offset+len(two) > len(one) {
			// no room in one[offset:] to contain two
			return "unequal"
		}
		if Equal(one[offset:offset+len(two)], two) {
			if len(one) == len(two) {
				return "equal"
			}
			if flipped {
				return "sublist"
			}
			return "superlist"
		}
	}
	return "unequal"
}

func Equal(one, two []int) bool {
	if one == nil && two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
