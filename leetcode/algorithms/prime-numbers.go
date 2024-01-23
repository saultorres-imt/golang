package algorithms

import (
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sumDigits(num int) int {
	suma := 0
	for num > 0 {
		suma += num % 10
		num /= 10
	}
	return suma
}

func SumPrimeNumbers(inicio, fin int) []int {
	var primos []int
	for i := inicio; i <= fin; i++ {
		if isPrime(i) && sumDigits(i) >= 10 {
			primos = append(primos, i)
		}
	}
	return primos
}
