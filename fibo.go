package main

import "fmt"

func fibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	n := 5
	fmt.Println(fibo(n))
}
