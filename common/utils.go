// Common functions and utilities
package common

import (
	"fmt"
)

const MAX_TRAVERSE_LENGTH = 50
const MAX_RANDOM_VALUE = 1000

// Collects input from stdin, and ensures it is an integer
func GetIntInput() int {
	var input int
	fmt.Scanf("%d", &input)

	return input
}

// Print a menu with all the bells and whistles
func PrintMenu(ops string) {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("Select an operation to perform:")
	fmt.Println(ops)
	fmt.Println("---------------------------------------------")
}
