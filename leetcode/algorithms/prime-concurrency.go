package algorithms

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func isPrimeC(num int) bool {
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

func sumDigitsC(num int) int {
	suma := 0
	for num > 0 {
		suma += num % 10
		num /= 10
	}
	return suma
}

func SumPrimeNumbersC(start, end int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if isPrimeC(i) && !(sumDigitsC(i) >= 10) {
			c <- i
		}
	}
}

func ModelMain(n1, n2 int) []int {
	var primes []int
	var input string
	reader := bufio.NewReader(os.Stdin)
	c := make(chan int)
	var waitGroup sync.WaitGroup

	fmt.Println("How many goroutines do you want?")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inicio := time.Now()
	input = strings.TrimSpace(input)

	inumRoutines, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	routineRange := (n1 - n2) / inumRoutines

	for i := 0; i < inumRoutines; i++ {
		waitGroup.Add(1)
		start := n1 + i*routineRange
		end := start + routineRange - 1
		if i == inumRoutines-1 {
			end = n2
		}
		go SumPrimeNumbersC(start, end, c, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
		close(c)
	}()

	for prime := range c {
		primes = append(primes, prime)
	}
	duracion := time.Since(inicio) // Calcular la duración
	fmt.Printf("Tiempo de ejecución: %s\n", duracion)
	return primes
}
