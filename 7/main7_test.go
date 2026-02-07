package main

import (
	"testing"
	"time"
)

func TestRandomStream(t *testing.T) {
	t.Run("should generate random numbers within range", func(t *testing.T) {
		stop := make(chan struct{})
		ch := randomStream(100, stop)

		for i := 0; i < 5; i++ {
			select {
			case val := <-ch:
				if val < 0 || val >= 100 {
					t.Errorf("value %d out of range [0, 100)", val)
				}
			case <-time.After(100 * time.Millisecond):
				t.Fatal("timeout reading from channel")
			}
		}

		close(stop)
	})

	t.Run("should stop when stop channel is closed", func(t *testing.T) {
		stop := make(chan struct{})
		ch := randomStream(50, stop)

		close(stop)

		time.Sleep(10 * time.Millisecond)

		select {
		case _, ok := <-ch:
			if ok {
				t.Error("channel should be closed after stop")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("channel should be readable after stop")
		}
	})

	t.Run("should handle stop signal after some values", func(t *testing.T) {
		stop := make(chan struct{})
		ch := randomStream(10, stop)

		for i := 0; i < 3; i++ {
			<-ch
		}

		close(stop)

		time.Sleep(10 * time.Millisecond)

		_, ok := <-ch
		if ok {
			t.Error("channel should be closed after stop signal")
		}
	})

	t.Run("should work with number=1 (always 0)", func(t *testing.T) {
		stop := make(chan struct{})
		ch := randomStream(1, stop)

		for i := 0; i < 3; i++ {
			select {
			case val := <-ch:
				if val != 0 {
					t.Errorf("expected 0, got %d", val)
				}
			case <-time.After(100 * time.Millisecond):
				t.Fatal("timeout reading from channel")
			}
		}

		close(stop)
	})
}
