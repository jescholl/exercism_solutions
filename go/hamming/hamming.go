package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return -1, errors.New("Length mismatch")
	}

	for i, _ := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
