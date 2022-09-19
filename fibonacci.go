package main

func fibonacci(n int) int {
	fibPrev := 0
	fib := 1
	var temp int
	for i := 2; i <= n; i++ {
		temp = fib
		fib = fibPrev + fib
		fibPrev = temp
	}

	return fib
}
