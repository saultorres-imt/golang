package main

import "fmt"

// Main function tu run all leet code programming challenges
func main() {
	fmt.Println("Welcome")

	// Can Place Flowers
	flowerbed := []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0}
	fmt.Println(canPlaceFlowers(flowerbed, 17))

	// Max Palindrome substring
	longString := " a b c d ed c bt t bc d ea"
	fmt.Println(longesPalindromeSubs(longString))
}
