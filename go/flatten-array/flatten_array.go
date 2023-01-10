package flatten

func Flatten(nested interface{}) []interface{} {
	output := make([]interface{}, 0)
	loop(nested, output)
	return output
}

func loop(t interface{}, output []interface{}) {

	switch t := t.(type) {
	case int:
		output = append(output, t)
	case []int:
		for _, v := range t {
			output = append(output, v)
		}
	case []interface{}:
		loop(t, output)
	}
}
