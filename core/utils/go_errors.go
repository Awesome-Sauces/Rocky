package utils

import "log"

// Simple error checking
func Check(exception error) {

	// If error is not nil (Means an error
	// was passed) will throw error

	if exception != nil {
		// Uses log to log the error
		log.Fatal(exception)
	}
}
