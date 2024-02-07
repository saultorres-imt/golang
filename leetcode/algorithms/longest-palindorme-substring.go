package algorithms

import (
	"fmt"
	"strings"
)

func LongestPalindromeSubs(str string) string {
	var palindrome string
	strNoSpaces := strings.ReplaceAll(str, " ", "")
	lowerCaseStr := strings.ToLower(strNoSpaces)
	lastPosition := len(lowerCaseStr) - 1

	for i := 0; i < lastPosition-1; i++ {
		if len(palindrome) > (lastPosition - i) {
			return palindrome
		}
		for z := lastPosition; z > i+1; z-- {
			start := 0 + i
			for j := z; j > start+1; start, j = start+1, j-1 {
				if lowerCaseStr[start] == lowerCaseStr[j] {
					if start+1 == j-1 {
						palindromeFound := lowerCaseStr[i : z+1]
						fmt.Println("Palindrome Found -> ", palindromeFound)
						if len(palindromeFound) > len(palindrome) {
							palindrome = palindromeFound
						}
					}
					continue
				} else {
					break
				}
			}
		}
	}
	return palindrome
}
