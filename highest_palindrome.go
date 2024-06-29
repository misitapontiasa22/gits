package main

// func main() {
// 	s1 := "3943"
// 	k1 := 1
// 	fmt.Println(highestPalindrome(s1, k1)) // Output: "3993"

// 	s2 := "092282"
// 	k2 := 3
// 	fmt.Println(highestPalindrome(s2, k2)) // Output: "992299"

//	s3 := "abbcccd"
//	queries3 := []int32{1, 3, 9, 8}
//	result3 := weightedStrings(s3, queries3)
//	fmt.Println(result3) // Output: [Yes, Yes, Yes, No]
// }

func highestPalindrome(input string, k int) string {
	n := len(input)
	sBytes := []byte(input)
	changes := make([]int, n)
	left, right := 0, n-1

	// First pass: Make the string a palindrome
	for left < right {
		if sBytes[left] != sBytes[right] {
			if sBytes[left] > sBytes[right] {
				sBytes[right] = sBytes[left]
			} else {
				sBytes[left] = sBytes[right]
			}
			changes[left], changes[right] = 1, 1
			k--
		}
		left++
		right--
	}

	// If we used more than k changes, it's not possible
	if k < 0 {
		return "-1"
	}

	left, right = 0, n-1

	// Second pass: Maximize the palindrome
	for left <= right {
		if left == right && k > 0 {
			sBytes[left] = '9'
		}
		if sBytes[left] < '9' {
			if changes[left] == 1 && k >= 1 {
				sBytes[left], sBytes[right] = '9', '9'
				k--
			} else if changes[left] == 0 && k >= 2 {
				sBytes[left], sBytes[right] = '9', '9'
				k -= 2
			}
		}
		left++
		right--
	}

	return string(sBytes)
}
