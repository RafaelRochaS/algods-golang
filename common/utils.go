// Common functions and utilities
package common

import "fmt"

// Collects input from stdin, and ensures it is an integer
func GetIntInput() int {
	var input int
	fmt.Scanf("%d", &input)

	return input
}