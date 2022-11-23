package utils

import (
	"fmt"
	"os"
)

// Passes os.Args length
func getArgLength() int {
	return len(os.Args)
}

// Test Function
func atest() {
	for index, element := range os.Args {
		fmt.Printf("TEST -> Index: %d | Element: %s\n", index, element)
	}
}

// Returns boolean based on if the argument
// exists and is named the same or not
func IsArg(toGet string, arg int) bool {
	// If arg = 0 return false
	// We do this because 0 is the initial
	// command which will never change
	if arg == 0 {
		return false
	}

	// Checks if the values match
	// Yes = return true No = return false
	if os.Args[arg] == toGet {
		return true
	}

	// Default return value
	return false
}

// Returns boolean based on if the argument exists or not
func ArgExists(arg int) bool {
	// If arg is greater than len then we
	// return false, or if it is less than 1
	if len(os.Args) <= arg || arg < 1 {
		return false
	}

	// Default return value
	return true
}

// Returns argumeent name
func GetArgName(arg int) string {
	// If arg = 0 return empty
	// We do this because 0 is the initial
	// command which will never change
	if arg == 0 || !ArgExists(arg) {
		return ""
	}

	// Returns the request arg name
	return os.Args[arg]
}
