// cmd/cheetah-clash/main.go
package main

import (
	"fmt"

	"github.com/DeepakGunpal/stringReverser/pkg/stringutils"
)

func main() {
	fmt.Println("Cheetah Clash")

	// Example usage of stringReverse
	original := "Hello, World!"
	reversed := stringutils.Reverse(original)
	fmt.Printf("Original: %s\nReversed: %s\n", original, reversed)

	// Your main application logic goes here
}
