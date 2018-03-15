package acronym

import "strings"

func Abbreviate(s string) string {
	var acronym string

	s = strings.Replace(s, "-", " ", -1)

	for _, word := range strings.Split(s, " ") {
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}
