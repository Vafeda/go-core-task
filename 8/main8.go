package main

import (
	"log"
	"time"
)

const (
	countGoroutine = 10
)

func main() {
	sem := NewSemaphore(2)

	for i := 0; i < countGoroutine; i++ {
		go func(numberGoroutine int, s *Semaphore) {
			s.Add()
			defer s.Done()

			log.Printf("Goroutine #%d is running", numberGoroutine)
			time.Sleep(1 * time.Second)
		}(i, sem)
	}

	sem.Wait()
}

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(count int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, count),
	}
}

func (s *Semaphore) Add() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Done() {
	<-s.ch
}

func (s *Semaphore) Wait() {
	for {
		if len(s.ch) == 0 {
			close(s.ch)
			return
		}
	}
}
