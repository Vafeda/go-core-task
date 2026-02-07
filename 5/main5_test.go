package main

import (
	"reflect"
	"testing"
)

func TestIntersects(t *testing.T) {
	tests := []struct {
		name          string
		slice1        []int
		slice2        []int
		wantBool      bool
		wantIntersect []int
	}{
		{
			name:          "basic intersection",
			slice1:        []int{65, 3, 58, 678, 64},
			slice2:        []int{64, 2, 3, 43},
			wantBool:      true,
			wantIntersect: []int{64, 3},
		},
		{
			name:          "no intersection",
			slice1:        []int{1, 2, 3},
			slice2:        []int{4, 5, 6},
			wantBool:      false,
			wantIntersect: []int{},
		},
		{
			name:          "all elements intersect",
			slice1:        []int{1, 2, 3},
			slice2:        []int{1, 2, 3},
			wantBool:      true,
			wantIntersect: []int{1, 2, 3},
		},
		{
			name:          "slice1 empty",
			slice1:        []int{},
			slice2:        []int{1, 2, 3},
			wantBool:      false,
			wantIntersect: []int{},
		},
		{
			name:          "slice2 empty",
			slice1:        []int{1, 2, 3},
			slice2:        []int{},
			wantBool:      false,
			wantIntersect: []int{},
		},
		{
			name:          "duplicates in both slices",
			slice1:        []int{1, 1, 2, 2, 3},
			slice2:        []int{1, 2, 2, 3, 3},
			wantBool:      true,
			wantIntersect: []int{1, 2, 3},
		},
		{
			name:          "negative numbers intersection",
			slice1:        []int{-1, -2, 0, 1},
			slice2:        []int{-2, 0, 2, -3},
			wantBool:      true,
			wantIntersect: []int{-2, 0},
		},
		{
			name:          "single element intersection",
			slice1:        []int{100, 200, 300},
			slice2:        []int{300, 400, 500},
			wantBool:      true,
			wantIntersect: []int{300},
		},
		{
			name:          "slice1 nil",
			slice1:        nil,
			slice2:        []int{1, 2, 3},
			wantBool:      false,
			wantIntersect: []int{},
		},
		{
			name:          "slice2 nil",
			slice1:        []int{1, 2, 3},
			slice2:        nil,
			wantBool:      false,
			wantIntersect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotIntersect := intersects(tt.slice1, tt.slice2)

			if gotBool != tt.wantBool {
				t.Errorf("intersects() bool = %v, want %v", gotBool, tt.wantBool)
			}

			if !reflect.DeepEqual(gotIntersect, tt.wantIntersect) {
				t.Errorf("intersects() intersect = %v, want %v", gotIntersect, tt.wantIntersect)
			}

			if len(gotIntersect) > 0 {
				seen := make(map[int]bool)
				for _, val := range gotIntersect {
					if seen[val] {
						t.Errorf("intersects() returned duplicate value %v in intersection", val)
					}
					seen[val] = true
				}
			}
		})
	}
}
