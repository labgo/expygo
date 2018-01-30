package main

import "fmt"

// A more generator-like style

func fib2(n int) <-chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a
		}
		close(c)
	}()
	return c
}

func main() {
	for x := range fib2(10) {
		fmt.Println(x)
	}
}
