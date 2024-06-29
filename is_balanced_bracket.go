package main

// func main() {
// 	expressions := []string{
// 		"(){}[]",       // Balanced
// 		"{[()]}",       // Balanced
// 		"{[(])}",       // Not balanced
// 		"{{[[(())]]}}", // Balanced
// 	}

// 	for _, expr := range expressions {
// 		if isBalanced(expr) {
// 			fmt.Printf("%s is balanced.\n", expr)
// 		} else {
// 			fmt.Printf("%s is not balanced.\n", expr)
// 		}
// 	}
// }

// Function to check if brackets are balanced
func isBalanced(expression string) bool {
	stack := make([]rune, 0)

	// Map to store matching brackets
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Iterate through each character in the expression
	for _, char := range expression {
		switch char {
		// If opening bracket, push to stack
		case '(', '{', '[':
			stack = append(stack, char)
		// If closing bracket, check stack for matching opening bracket
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop from stack
		}
	}

	// After iterating through expression, stack should be empty for balanced brackets
	return len(stack) == 0
}
