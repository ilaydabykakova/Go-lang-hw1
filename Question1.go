package main

import "fmt"

// Get all prime factors given numbers

func PrimeFactors(A, B, M int) (a int) {

	// Get the number of that divide n
	var result int
	sum := 0

	// Check to A if  numbers are divisible by i.
	for i := 1; i <= A; i++ {

		// If the numbers are divisible,then add it to sum
		if i%B == 0 || i%M == 0 {
			sum = sum + i
			a++
			fmt.Println("Number of divisible", i)
		}
	}

	fmt.Println("Counting numbers of divisible number", a)
	// Return the sum
	result = sum % a
	fmt.Println("Output", result)

	return sum
}

func main() {

	fmt.Println("Sum of divisible number", PrimeFactors(88, 8, 11))

}
