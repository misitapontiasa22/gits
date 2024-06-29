package main

import "fmt"

func main() {
	s := "abccddde"
	queries := []int32{1, 3, 12, 5, 9, 10}
	result := weightedStrings(s, queries)
	fmt.Println(result) // Output: [Yes Yes Yes Yes No No]
	
	s2 := "aabbcccde"
	queries2 := []int32{1, 3, 9, 8, 5}
	result2 := weightedStrings(s2, queries2)
	fmt.Println(result2) // Output: [Yes Yes Yes No Yes]

	s3 := "abbcccd"
	queries3 := []int32{1, 3, 9, 8}
	result3 := weightedStrings(s3, queries3)
	fmt.Println(result3) // Output: [Yes, Yes, Yes, No]
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
