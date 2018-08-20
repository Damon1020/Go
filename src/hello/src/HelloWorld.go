package main

import (
	//"cal"
	"fmt"
	"sync"
	"time"
)

func main() {
	//	x := cal.Add(1, 2)
	//	fmt.Println("x=", x)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Hi，I'm Damon!", n)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Hi，I'm Elena~", n)
			time.Sleep(1 * time.Second)
			//wg.Done()
			wg.Add(-1)
		}(i)
	}

	wg.Wait()
}
