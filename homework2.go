package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
)

func factorial(n int) {
	var res = 1
	var channel =make(chan int,20)
	for i := 1; i <= n; i++ {
		res *= i
	}
	channel <- res
	myres[n]=<-channel
}

func main() {

	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}