// cmd/cheetah-clash/main.go
package main

import (
	"fmt"
)

// stringReverse returns the reverse of the given string.
func stringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Cheetah Clash")

	// Example usage of stringReverse
	original := "Hello, World!"
	reversed := stringReverse(original)
	fmt.Printf("Original: %s\nReversed: %s\n", original, reversed)

	// Your main application logic goes here
}
