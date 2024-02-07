package algorithms

import (
	"fmt"
)

func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	var result []byte
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += cycleLen {
			result = append(result, s[j+i])
			fmt.Println("Result is :", string(result))
			if i != 0 && i != numRows-1 && j+cycleLen-i < len(s) {
				result = append(result, s[j+cycleLen-i])
				fmt.Println("Result is :", string(result))
			}
		}
	}

	return string(result)
}

/*
P     I    N
A   L S  I G
Y A   H R
P     I

0 - 6 - 12 - 1 - 5 - 7 - 11 - 13 - 2 - 4 - 8 - 10 - 3 - 9
*/
