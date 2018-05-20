package main

import (
	"fmt"
)

// By considering the terms in the Fibonacci sequence whose
// values do not exceed four million, find the sum of the even-valued terms.

// Fibonacci closure.
func fibonacci() func() uint32 {
	i := 0
	a := 0
	b := 1
	var e uint32
	return func() uint32 {
		i = a
		a = a + b
		b = i
		if i%2 == 0 {
			e += uint32(i)
		}
		return e
	}
}

// Fibonacci recursive.
func fib(n uint32) uint32 {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	f := fibonacci()
	var sum uint32
	for i := f(); i <= 4000000; i += f() {
		sum = i
	}
	fmt.Println(sum)
}
