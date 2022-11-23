package file

import (
	"os"

	// Rocky Stuff
	"github.com/Awesome-Sauces/Rocky/core/utils"
)

// Stores the file in here
var file_dat string

// Returns the file data POINTER
func GetFile() *string {
	return &file_dat
}

// Refresh the file data with provided string
func refreshFile(String *string) *string {
	// Stores the pointer of the provided string in such
	// that a copy of the string isn't made each
	// call of the function
	file_dat = *String
	return GetFile()
}

// Loads the file into memory and then returns it's pointer
func LoadFile(filename string) *string {
	var data, err = os.ReadFile(filename)

	// Error checks *err*
	utils.Check(err)

	// Creating temp var string() doesn't wanna work
	temp := string(data)

	// Just saves and returns the file data
	return refreshFile(&temp)
}
