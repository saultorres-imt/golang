package algorithms

import (
	"fmt"
	"strings"
)

func ReverseVowels(s string) string {
	var vowels []rune
	var result string

	for _, v := range s {
		switch v {
		case 'a':
			vowels = append(vowels, 'a')
		case 'e':
			vowels = append(vowels, 'e')
		case 'i':
			vowels = append(vowels, 'i')
		case 'o':
			vowels = append(vowels, 'o')
		case 'u':
			vowels = append(vowels, 'u')
		case 'A':
			vowels = append(vowels, 'A')
		case 'E':
			vowels = append(vowels, 'E')
		case 'I':
			vowels = append(vowels, 'I')
		case 'O':
			vowels = append(vowels, 'O')
		case 'U':
			vowels = append(vowels, 'U')
		default:
		}
	}
	j := len(vowels) - 1
	for _, v := range s {
		switch v {
		case 'a':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'e':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'i':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'o':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'u':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'A':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'E':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'I':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'O':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		case 'U':
			result = fmt.Sprintf("%s%s", result, string(vowels[j]))
			j--
		default:
			result = fmt.Sprintf("%s%s", result, string(v))
		}
	}
	return result
}

func reverseVowels(s string) string {
	low, high := 0, len(s)-1
	runeSlice := []rune(s)

	for low < high {
		for low < high && !isVowel(runeSlice[low]) {
			low++
		}
		for low < high && !isVowel(runeSlice[high]) {
			high--
		}

		if low < high {
			runeSlice[low], runeSlice[high] = runeSlice[high], runeSlice[low]
			low++
			high--
		}
	}

	return string(runeSlice)
}

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.Contains(vowels, string(r))
}
