package algorithms

import (
	"strings"
)

func GcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		if ok := strings.Contains(str1, str2); ok {
			if len(str1)%len(str2) == 0 {
				if ok = verifyString(str2, str1); ok {
					return str2
				} else {
					return ""
				}
			} else {
				gcd := findGCD(len(str1), len(str2))
				gcdStr := string(str1[:gcd])
				if ok = verifyString(gcdStr, str1); ok {
					if ok = verifyString(gcdStr, str2); ok {
						return gcdStr
					}
				}
			}
		} else {
			return ""
		}
	} else if len(str2) > len(str1) {
		if ok := strings.Contains(str2, str1); ok {
			if len(str1)/len(str2) == 2 {
				if ok = verifyString(str1, str2); ok {
					return str1
				} else {
					return ""
				}
			} else {
				gcd := findGCD(len(str2), len(str1))
				gcdStr := string(str2[:gcd])
				if ok = verifyString(gcdStr, str2); ok {
					if ok = verifyString(gcdStr, str1); ok {
						return gcdStr
					}
				}
			}
		}
	} else {
		if ok := strings.Contains(str1, str2); ok {
			if ok = verifyString(str1, str2); ok {
				return str1
			} else {
				return ""
			}
		} else {
			return ""
		}
	}
	return ""
}

func verifyString(gcd string, str string) bool {
	reps := len(str) / len(gcd)
	var start int

	for i := 0; i < reps; i++ {
		if ok := strings.Contains(string(str[start:start+len(gcd)]), gcd); ok {
			start += len(gcd)
			continue
		} else {
			return false
		}
	}
	return true
}

func findGCD(n int, m int) int {
	for m != 0 {
		t := m
		m = n % m
		n = t
	}
	return n
}

// ABCABC ABABAB
func gcdOfStrings(str1, str2 string) string {
	lengthGCD := findGcd(len(str1), len(str2))

	candidate := str1[:lengthGCD]
	if strings.Repeat(candidate, len(str1)/lengthGCD) == str1 && strings.Repeat(candidate, len(str2)/lengthGCD) == str2 {
		return candidate
	}
	return ""
}

func findGcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
