package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	length := len(runes)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	output = string(runes)

	return output
}
