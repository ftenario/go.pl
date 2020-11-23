package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	//Squarer
	go func() {
		for {
			c := <-naturals
			squares <- c * c
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
