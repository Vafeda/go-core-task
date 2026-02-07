package main

import (
	"sync"
	"testing"
	"time"
)

func TestNewSemaphore(t *testing.T) {
	capacity := 3
	sem := NewSemaphore(capacity)

	if cap(sem.ch) != capacity {
		t.Errorf("Expected capacity %d, got %d", capacity, cap(sem.ch))
	}

	if len(sem.ch) != 0 {
		t.Errorf("Expected empty semaphore, got %d items", len(sem.ch))
	}
}

func TestSemaphore_AddDone(t *testing.T) {
	capacity := 2
	sem := NewSemaphore(capacity)

	sem.Add()
	sem.Add()

	if len(sem.ch) != capacity {
		t.Errorf("Expected full semaphore (%d items), got %d", capacity, len(sem.ch))
	}

	sem.Done()
	sem.Done()

	if len(sem.ch) != 0 {
		t.Errorf("Expected empty semaphore, got %d items", len(sem.ch))
	}
}

func TestSemaphore_ConcurrentAccess(t *testing.T) {
	capacity := 3
	sem := NewSemaphore(capacity)
	var wg sync.WaitGroup

	for i := 0; i < capacity*2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Add()
			defer sem.Done()

			// Имитация работы
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	wg.Wait()

	if len(sem.ch) != 0 {
		t.Errorf("Semaphore should be empty after all goroutines finished, got %d items", len(sem.ch))
	}
}

func TestSemaphore_Wait(t *testing.T) {
	sem := NewSemaphore(2)

	sem.Add()
	sem.Add()

	done := make(chan bool)
	go func() {
		sem.Wait()
		done <- true
	}()

	select {
	case <-done:
		t.Error("Wait should not complete when semaphore is not empty")
	case <-time.After(100 * time.Millisecond):
	}

	sem.Done()
	sem.Done()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait should complete when semaphore is empty")
	}
}
