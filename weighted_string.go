package main

import "fmt"

func main() {
	s := "abccddde"
	queries := []int32{1, 3, 12, 5, 9, 10}
	result := weightedStrings(s, queries)
	fmt.Println(result) // Output: [Yes Yes Yes Yes No No]
}

func weightedStrings(s string, queries []int32) []string {
	weights := make(map[int32]bool)
	var currentWeight int32
	var prev rune

	for _, char := range s {
		if char == prev {
			currentWeight += int32(char - 'a' + 1)
		} else {
			currentWeight = int32(char - 'a' + 1)
		}
		weights[currentWeight] = true
		prev = char
	}

	results := make([]string, len(queries))
	for i, q := range queries {
		if weights[q] {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	return results
}
