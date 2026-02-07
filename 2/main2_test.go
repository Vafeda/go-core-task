package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "all even numbers",
			input:    []int{2, 4, 6, 8},
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "all odd numbers",
			input:    []int{1, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "mixed numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "with zero",
			input:    []int{0, 1, 2, 3},
			expected: []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)

			if len(got) != len(tt.expected) {
				t.Errorf("sliceExample() length = %v, want %v", len(got), len(tt.expected))
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("sliceExample()[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}

			if len(tt.input) > 0 {
				originalLen := len(tt.input)
				_ = sliceExample(tt.input)
				if len(tt.input) != originalLen {
					t.Errorf("original slice was modified")
				}
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		number    int
		wantLen   int
		shouldErr bool
	}{
		{
			name:    "add to empty slice",
			input:   []int{},
			number:  5,
			wantLen: 2,
		},
		{
			name:    "add to non-empty slice",
			input:   []int{1, 2, 3},
			number:  4,
			wantLen: 5,
		},
		{
			name:    "add negative number",
			input:   []int{1, 2},
			number:  -1,
			wantLen: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addElements(tt.input, tt.number)

			if tt.shouldErr && err == nil {
				t.Errorf("addElements() expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("addElements() unexpected error: %v", err)
			}

			if !tt.shouldErr {
				if len(got) != tt.wantLen {
					t.Errorf("addElements() length = %v, want %v", len(got), tt.wantLen)
				}

				for i := 0; i < len(tt.input); i++ {
					if got[i] != tt.input[i] {
						t.Errorf("addElements()[%d] = %v, want %v (original element)", i, got[i], tt.input[i])
					}
				}

				originalLen := len(tt.input)
				_, _ = addElements(tt.input, tt.number)
				if len(tt.input) != originalLen {
					t.Errorf("original slice was modified")
				}
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		shouldErr bool
	}{
		{
			name:  "empty slice",
			input: []int{},
		},
		{
			name:  "single element",
			input: []int{1},
		},
		{
			name:  "multiple elements",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with negative numbers",
			input: []int{-1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := copySlice(tt.input)

			if tt.shouldErr && err == nil {
				t.Errorf("copySlice() expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("copySlice() unexpected error: %v", err)
			}

			if !tt.shouldErr {
				if len(got) != len(tt.input) {
					t.Errorf("copySlice() length = %v, want %v", len(got), len(tt.input))
				}

				for i := range got {
					if got[i] != tt.input[i] {
						t.Errorf("copySlice()[%d] = %v, want %v", i, got[i], tt.input[i])
					}
				}

				if len(tt.input) > 0 {
					originalFirst := tt.input[0]
					copyFirst := got[0]

					tt.input[0] = originalFirst + 100

					if got[0] != copyFirst {
						t.Errorf("copy was affected by modifying original slice")
					}

					tt.input[0] = originalFirst
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		index       int
		expected    []int
		shouldPanic bool
	}{
		{
			name:     "remove first element",
			input:    []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "remove middle element",
			input:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "remove last element",
			input:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "remove from single element slice",
			input:    []int{1},
			index:    0,
			expected: []int{},
		},
		{
			name:     "remove from empty slice",
			input:    []int{},
			index:    0,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.shouldPanic {
					t.Errorf("removeElement() panicked unexpectedly: %v", r)
				}
			}()

			got := removeElement(tt.input, tt.index)

			if len(got) != len(tt.expected) {
				t.Errorf("removeElement() length = %v, want %v", len(got), len(tt.expected))
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("removeElement()[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}

			if len(tt.input) > 0 {
				originalLen := len(tt.input)
				_ = removeElement(tt.input, tt.index)
				if len(tt.input) != originalLen {
					t.Errorf("original slice was modified")
				}
			}
		})
	}
}
