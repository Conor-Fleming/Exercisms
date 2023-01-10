package flatten

func Flatten(nested interface{}) []interface{} {
	output := make([]interface{}, 0)
	result := nested.([]interface{})

	for _, val := range result {
		switch val.(type) {
		case nil:
			continue
		case []interface{}:
			output = append(output, Flatten(val)...)
		default:
			output = append(output, val)
		}
	}

	return output
}
