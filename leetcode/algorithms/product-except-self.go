package algorithms

import "fmt"

// [1,2,3,4]
func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	fmt.Println("nums :", nums)

	for i := 1; i < len(nums); i++ {
		answer[i] = nums[i-1] * answer[i-1]
		fmt.Println(answer)
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	fmt.Println(answer)

	return answer
}
