package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var a = make(chan struct{}, 1)
	var b = make(chan struct{}, 1)

	b <- struct{}{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 100; i += 2 {
			<-b
			fmt.Printf("%d\n", i)
			a <- struct{}{}
		}

	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-a
			fmt.Printf("%d\n", i)
			b <- struct{}{}
		}

	}()
	wg.Wait()
}
