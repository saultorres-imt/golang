package algorithms

import (
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
	/*
	   noSpaces := regexp.MustCompile(`\s+`)
	   wordNoSpaces := noSpaces.ReplaceAllString(strings.TrimSpace(s), " ")
	   split := strings.Split(wordNoSpaces, " ")

	   	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
	   		split[i], split[j] = split[j], split[i]
	   	}

	   result = strings.Join(split[:], " ")
	   return result
	*/
}
