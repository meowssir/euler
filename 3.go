package main

import (
	"fmt"
)

func main() {
	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143?
	n := 600851475143
	i := 2
	for i*i < n {
		for n%i == 0 {
			n = n / i
		}
		i = i + 1
	}
	fmt.Println(n)
}
