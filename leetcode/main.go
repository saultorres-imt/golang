package main

import (
	"fmt"
	"leetcode/algorithms"
)

// Main function tu run all leetcode programming challenges
func main() {
	fmt.Println("Welcome")

	// Can Place Flowers
	// flowerbed := []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0}
	// fmt.Println(algorithms.CanPlaceFlowers(flowerbed, 17))

	// Max Palindrome substring
	// longString := "a bcb ebc anita la va la tinat a"
	// fmt.Println("Longest palindrome found -> ", algorithms.LongestPalindromeSubs(longString))

	// Merge Strings Alternately
	word1, word2 := "abc", "pqrft"
	fmt.Println("The merge string would be -> ", algorithms.MergeAlternately(word1, word2))
}
