package algorithms

import "fmt"

func KidsWithCandies(candies []int, extraCandies int) []bool {
	var greatestNumber int
	result := make([]bool, len(candies))

	for _, v := range candies {
		if v > greatestNumber {
			greatestNumber = v
		}
	}

	fmt.Println("Greates Number is -> ", greatestNumber)

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= greatestNumber {
			result[i] = true
		}
	}

	return result
}
