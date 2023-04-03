package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	// Rocky imports
	"github.com/Awesome-Sauces/Rocky/core/eval"
	"github.com/Awesome-Sauces/Rocky/core/file"
	"github.com/Awesome-Sauces/Rocky/core/lexer"
	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
	"github.com/Awesome-Sauces/Rocky/core/utils"
	"github.com/novalagung/golpal"
)

func main() {

	// Example: -tokenize enumNumber
	if utils.IsArg("-tokenize", 1) && utils.ArgExists(2) {
		i, err := strconv.Atoi(utils.GetArgName(2))

		utils.Check(err)

		token := tokenizer.CreateToken(tokenizer.TOKEN(i), "{", 0)

		fmt.Println(token.GetType().ToString())
	}

	// Example: -eval filename.extension
	if utils.IsArg("-eval", 1) && utils.ArgExists(2) {

		file.LoadFile(utils.GetArgName(2))

		fmt.Printf(*file.GetFile())
	}

	if utils.IsArg("-testGoEval", 1) && utils.ArgExists(2) {
		file.LoadFile(utils.GetArgName(2))

		temp := *file.GetFile()

		output, err := golpal.New().ExecuteRaw(temp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)

	}

	// Example: -test filename.extension
	if utils.IsArg("-test", 1) && utils.ArgExists(2) {
		file.LoadFile(utils.GetArgName(2))
		//fmt.Printf("%s\n", *file.GetFile())
		tokens := lexer.LoopString(file.GetFile())
		writeTokenMapToFile(tokens, "rocky-lexOut.json")

		eval.Eval(tokens)

		/*
			for _, element := range tokens {
				fmt.Println("TYPE -> " + element.GetType().ToString() + " -> VALUE -> " + element.GetValue() + " -> ORDER -> " + strconv.Itoa(element.GetOrder()))
			}
		*/

	} else if utils.IsArg("-test", 1) && !utils.ArgExists(2) {
		fmt.Printf("Error: Please enter filename!")
		os.Exit(1)
	}

	fmt.Println()
}

// writeTokenMapToFile takes a map with integer keys and Token values, marshals the data to JSON,
// and writes it to a file with the given filename.
func writeTokenMapToFile(tokenMap map[int]*tokenizer.Token, filename string) error {
	// Create a new slice to hold the marshaled Token values
	marshaledTokens := make([]map[string]interface{}, len(tokenMap))

	// Loop over the Token map and marshal each Token value to JSON
	for i, token := range tokenMap {
		marshaledTokens[i] = map[string]interface{}{
			"Type":  token.Type.ToString(),
			"Value": token.Value,
			"Order": token.Order,
		}
	}

	// Marshal the marshaledTokens slice to JSON
	jsonData, err := json.Marshal(marshaledTokens)
	if err != nil {
		return err
	}

	// Write the JSON data to a file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
