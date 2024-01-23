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
	// word1, word2 := "abc", "pqrft"
	// fmt.Println("The merge string would be -> ", algorithms.MergeAlternately(word1, word2))

	// Greates Common Divisor os Strings
	// str1, str2 := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	// fmt.Println("Greates common divisor is -> ", algorithms.GcdOfStrings(str1, str2))

	// Kids with candies
	// candies, extraCandies := []int{2, 3, 5, 1, 3}, 3
	// fmt.Println("Result -> ", algorithms.KidsWithCandies(candies, extraCandies))

	// Reverse Vowels
	// str := "e0P"
	// fmt.Println("Result -> ", algorithms.ReverseVowels(str))

	// Reverse Strings
	// str := "  hello  World "
	// fmt.Println("Result -> ", algorithms.ReverseWords(str))

	// Insert Delete GetRandom
	// str := "  hello  World "
	// fmt.Println("Result -> ", algorithms.ReverseWords(str))

	// obj := algorithms.Constructor()
	// fmt.Println(obj.Insert(1))
	// fmt.Println(obj.Remove(1))
	// fmt.Println(obj.GetRandom())

	// Product Except Self
	// iArr := []int{1, 2, 3, 4}
	// fmt.Println("The array is :", algorithms.ProductExceptSelf(iArr))

	// Prime Numbers
	fmt.Println("Prime numbers that the sum of its digits is 10 or more :", algorithms.ModelMain(1, 50))
}
