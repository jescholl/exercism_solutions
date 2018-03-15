package reverse

func String(phrase string) string {
	var output []byte

	for i := len(phrase) - 1; i >= 0; i-- {
		output = append(output, phrase[i])
	}

	return string(output)
}
