package main

import "fmt"

func get_prime(n int) {
	o, wait := make(chan int), make(chan struct{})
	go Prime(o, wait)
	for i := 2; i <= n; i++ {
		o<-i
	}
	close(o)
	<-wait
}

func Prime(seq chan int, wait chan struct{}) {
	prime, ok := <-seq
	if !ok {
		close(wait)
		return
	}
	fmt.Println(prime)
	out := make(chan int)
	go Prime(out, wait)
	for num := range seq {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out)
}
func main() {
	get_prime(10000)
}