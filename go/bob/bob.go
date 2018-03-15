package bob

import "regexp"

func Hey(remark string) string {

	var question = regexp.MustCompile(`\?\s*$`)
	var yell = regexp.MustCompile(`[A-Z]{2,}`)
	var calm = regexp.MustCompile(`[a-z]{2,}`)
	var nothing = regexp.MustCompile(`^\s*$`)

	// ask a question
	if (!yell.MatchString(remark) || calm.MatchString(remark)) && question.MatchString(remark) {
		return "Sure."
	}
	// yell a question
	if question.MatchString(remark) && yell.MatchString(remark) {
		return "Calm down, I know what I'm doing!"
	}
	// yell
	if yell.MatchString(remark) && !calm.MatchString(remark) {
		return "Whoa, chill out!"
	}
	// nothing
	if nothing.MatchString(remark) {
		return "Fine. Be that way!"
	}
	// everything else
	return "Whatever."
}
