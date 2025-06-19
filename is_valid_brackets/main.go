package main

import "fmt"

func isValid(pattern string, m map[rune]rune) bool {
	stack := []rune{}

	for _, ch := range pattern {
		switch ch {
		case '(':
			stack = append(stack, ch)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != m[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	patterns := []string{"{[(())]}", "()()()()", ")()(", "{[()}]", "()(", ")()"}

	for _, pattern := range patterns {
		if isValid(pattern, map[rune]rune{
			')': '(',
			']': '[',
			'}': '{',
		}) {
			fmt.Println("--> Valid: ", pattern)
		} else {
			fmt.Println("--> Not Valid: ", pattern)
		}
	}

}
