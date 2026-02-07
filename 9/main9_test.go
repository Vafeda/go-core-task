package main

import (
	"sync"
	"testing"
)

func TestGenerateNumbers(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"Zero numbers", 0},
		{"One number", 1},
		{"Multiple numbers", 5},
		{"Many numbers", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan uint8, 10)
			wg := &sync.WaitGroup{}
			wg.Add(1)

			go GenerateNumbers(tt.count, ch, wg)
			wg.Wait()

			var count int
			for range ch {
				count++
			}

			if count != tt.count {
				t.Errorf("GenerateNumbers() generated %d numbers, want %d", count, tt.count)
			}
		})
	}
}

func TestProcessChannels(t *testing.T) {
	tests := []struct {
		name   string
		input  []uint8
		output []float64
	}{
		{
			name:   "Empty channel",
			input:  []uint8{},
			output: []float64{},
		},
		{
			name:   "Single value",
			input:  []uint8{2},
			output: []float64{8},
		},
		{
			name:   "Multiple values",
			input:  []uint8{0, 1, 3, 5},
			output: []float64{0, 1, 27, 125},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inChan := make(chan uint8, len(tt.input))
			outChan := make(chan float64, len(tt.output))
			wg := &sync.WaitGroup{}
			wg.Add(1)

			for _, v := range tt.input {
				inChan <- v
			}
			close(inChan)

			go ProcessChannels(inChan, outChan, wg)
			wg.Wait()

			var results []float64
			for v := range outChan {
				results = append(results, v)
			}

			if len(results) != len(tt.output) {
				t.Errorf("ProcessChannels() produced %d results, want %d", len(results), len(tt.output))
				return
			}

			for i, v := range results {
				if v != tt.output[i] {
					t.Errorf("ProcessChannels()[%d] = %v, want %v", i, v, tt.output[i])
				}
			}
		})
	}
}

func TestProcessChannels_CubeCalculation(t *testing.T) {
	testCases := []struct {
		input    uint8
		expected float64
	}{
		{0, 0},
		{1, 1}, 
		{2, 8},  
		{3, 27},  
		{10, 1000},
	}

	for _, tc := range testCases {
		t.Run("Calculation", func(t *testing.T) {
			inChan := make(chan uint8, 1)
			outChan := make(chan float64, 1)
			wg := &sync.WaitGroup{}
			wg.Add(1)

			inChan <- tc.input
			close(inChan)

			go ProcessChannels(inChan, outChan, wg)
			wg.Wait()

			result := <-outChan

			if result != tc.expected {
				t.Errorf("ProcessChannels(%d) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
