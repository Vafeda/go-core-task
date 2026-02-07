package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

const (
	channelCapacity = 5
)

func ProcessChannels(inChan <-chan uint8, outChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(outChan)

	for v := range inChan {
		outChan <- math.Pow(float64(v), 3)
	}
}

func PrintResults(results <-chan float64, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range results {
		fmt.Println(v)
	}
}

func GenerateNumbers(count int, numbersChan chan<- uint8, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(numbersChan)

	for i := 0; i < count; i++ {
		numbersChan <- uint8(rand.Uint32() % 100)
	}
}

func main() {
	ch1 := make(chan uint8, channelCapacity)
	ch2 := make(chan float64, channelCapacity)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go GenerateNumbers(6, ch1, wg)

	wg.Add(1)
	go ProcessChannels(ch1, ch2, wg)

	wg.Add(1)
	go PrintResults(ch2, wg)

	wg.Wait()
}
