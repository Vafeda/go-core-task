package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "basic difference",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "slice1 empty",
			slice1:   []string{},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "slice2 empty",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "both slices empty",
			slice1:   []string{},
			slice2:   []string{},
			expected: []string{},
		},
		{
			name:     "slice1 nil",
			slice1:   nil,
			slice2:   []string{"a", "b"},
			expected: []string{},
		},
		{
			name:     "slice2 nil",
			slice1:   []string{"a", "b"},
			slice2:   nil,
			expected: []string{"a", "b"},
		},
		{
			name:     "duplicates in slice1",
			slice1:   []string{"a", "a", "b", "b", "c"},
			slice2:   []string{"b", "c"},
			expected: []string{"a", "a"},
		},
		{
			name:     "all elements same",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "with empty strings",
			slice1:   []string{"", "a", "b", ""},
			slice2:   []string{"a", ""},
			expected: []string{"b"},
		},
		{
			name:     "case sensitive comparison",
			slice1:   []string{"Apple", "banana", "Cherry"},
			slice2:   []string{"apple", "BANANA", "cherry"},
			expected: []string{"Apple", "banana", "Cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := difference(tt.slice1, tt.slice2)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("difference() = %v, want %v", result, tt.expected)
			}

			if tt.slice1 != nil {
				originalLen := len(tt.slice1)
				_ = difference(tt.slice1, tt.slice2)
				if len(tt.slice1) != originalLen {
					t.Errorf("slice1 was modified")
				}
			}
		})
	}
}
