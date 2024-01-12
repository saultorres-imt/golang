package algorithms

func MergeAlternately(word1 string, word2 string) string {
	var mergeString string

	if len(word1) > len(word2) {
		for i := 0; i < len(word2); i++ {
			mergeString += string(word1[i]) + string(word2[i])
		}

		mergeString += string(word1[len(word2):])

	} else if len(word2) > len(word1) {
		for i := 0; i < len(word1); i++ {
			mergeString += string(word1[i]) + string(word2[i])
		}

		mergeString += string(word2[len(word1):])

	} else {
		for i := 0; i < len(word1); i++ {
			mergeString += string(word1[i]) + string(word2[i])
		}
	}

	return mergeString
}
