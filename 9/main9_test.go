package main

import (
	"sync"
	"testing"
)

// Тест для функции GenerateNumbers
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

			// Проверяем количество сгенерированных чисел
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

// Тест для функции ProcessChannels
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
			output: []float64{8}, // 2^3 = 8
		},
		{
			name:   "Multiple values",
			input:  []uint8{0, 1, 3, 5},
			output: []float64{0, 1, 27, 125}, // 0^3=0, 1^3=1, 3^3=27, 5^3=125
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inChan := make(chan uint8, len(tt.input))
			outChan := make(chan float64, len(tt.output))
			wg := &sync.WaitGroup{}
			wg.Add(1)

			// Заполняем входной канал
			for _, v := range tt.input {
				inChan <- v
			}
			close(inChan)

			// Запускаем обработку
			go ProcessChannels(inChan, outChan, wg)
			wg.Wait()

			// Собираем результаты
			var results []float64
			for v := range outChan {
				results = append(results, v)
			}

			// Проверяем результаты
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

// Тест для функции ProcessChannels с проверкой возведения в куб
func TestProcessChannels_CubeCalculation(t *testing.T) {
	testCases := []struct {
		input    uint8
		expected float64
	}{
		{0, 0},     // 0^3 = 0
		{1, 1},     // 1^3 = 1
		{2, 8},     // 2^3 = 8
		{3, 27},    // 3^3 = 27
		{10, 1000}, // 10^3 = 1000
	}

	for _, tc := range testCases {
		t.Run("Calculation", func(t *testing.T) {
			inChan := make(chan uint8, 1)
			outChan := make(chan float64, 1)
			wg := &sync.WaitGroup{}
			wg.Add(1)

			// Отправляем одно значение
			inChan <- tc.input
			close(inChan)

			// Запускаем обработку
			go ProcessChannels(inChan, outChan, wg)
			wg.Wait()

			// Получаем результат
			result := <-outChan

			if result != tc.expected {
				t.Errorf("ProcessChannels(%d) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
