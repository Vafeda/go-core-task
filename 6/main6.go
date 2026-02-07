package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	stopChan := make(chan struct{})
	randNum := randomStream(100, stopChan)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(<-randNum)
	}

	stopChan <- struct{}{}
	close(stopChan)
}

func randomStream(number int, stop <-chan struct{}) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for {
			select {
			case ch <- rand.Intn(number):
			case <-stop:
				return
			}
		}
	}()

	return ch
}
