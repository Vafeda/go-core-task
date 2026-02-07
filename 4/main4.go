package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(difference(slice1, slice2))
}

func difference(slice1 []string, slice2 []string) []string {
	if len(slice1) == 0 {
		return []string{}
	}

	if len(slice2) == 0 {
		return slice1
	}

	set := make(map[string]struct{}, len(slice2))
	for _, value := range slice2 {
		set[value] = struct{}{}
	}

	result := make([]string, 0, len(slice1))
	for _, value := range slice1 {
		if _, ok := set[value]; !ok {
			result = append(result, value)
		}
	}

	return result
}
