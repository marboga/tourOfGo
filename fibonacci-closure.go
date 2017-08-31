package main

import "fmt"

func main() {
	fib := Fibonacci()
	for i := 0; i < 90; i++ {
		fmt.Println(fib())
	}
}

func Fibonacci() func() int {
	var prevFib int
	fib := 1
	return func() int {
		newFib := prevFib + fib
		prevFib = fib
		fib = newFib
		return newFib
	}
}
