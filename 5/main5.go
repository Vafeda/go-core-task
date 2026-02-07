package main

import (
	"fmt"
)

func main() {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}
	fmt.Println(intersects(slice1, slice2))
}

func intersects(slice1 []int, slice2 []int) (bool, []int) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return false, []int{}
	}

	ma := make(map[int]int, len(slice1))
	for _, value := range slice1 {
		if _, ok := ma[value]; !ok {
			ma[value] = 1
		}
	}

	result := make([]int, 0, len(slice2))
	for _, value := range slice2 {
		if v, ok := ma[value]; ok && v == 1 {
			ma[value]++
			result = append(result, value)
		}
	}

	if len(result) == 0 {
		return false, []int{}
	}

	return true, result
}
