package main

import "fmt"

func ValidateInput(s string) (rune, error) {
	for _, char := range s {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("unsupported character: %c", char)
		}
	}
	return 0, nil
}