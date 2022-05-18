package reverse_string

func ReverseString(input string) (output string) {
	for _, s := range input {
		output = string(s) + output
	}

	return output
}
