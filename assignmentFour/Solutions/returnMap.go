package Solutions

func CountCharacters(word string) map[string]int {
	charCount := make(map[string]int)

	for i := 0; i < len(word); i++ {
		char := string(word[i])
		charCount[char]++
	}

	return charCount
}
