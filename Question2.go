package main

import "fmt"

func PrimeFactors(A, B, M int) (a int) {

	var result int
	sum := 0

	for i := 1; i <= A; i++ {

		if i%B == 0 || i%M == 0 {
			sum = sum + i
			a++
			fmt.Println("Number of divisible", i)
		}
	}

	fmt.Println("Counting numbers of divisible number", a)

	result = sum % a
	fmt.Println("Output", result)

	return sum
}

func main() {

	fmt.Println("Sum of divisible number", PrimeFactors(100, 7, 19))

}
