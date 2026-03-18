package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("send to closed channel  ", r)
		}
	}()
	res := make(chan int, 2)
	res <- 1
	res <- 2
	close(res)
	v, ok := <-res
	fmt.Println(v, ok)
	fmt.Println(<-res)
	fmt.Println(<-res)
}
