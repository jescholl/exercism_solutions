package flatten

// WORKING
func Flatten(in interface{}) []interface{} {

	out := make([]interface{}, 0)

	if inArr, isArray := in.([]interface{}); isArray {
		for _, item := range inArr {
			for _, f := range Flatten(item) {
				if f == nil {
					continue
				}
				out = append(out, f)
			}
		}
	} else {
		out = []interface{}{in}
	}

	return out
}
