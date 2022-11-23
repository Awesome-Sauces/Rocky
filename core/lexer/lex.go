package lexer

import (
	"fmt"
	"strings"

	// Rocky Imports
	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
)

func max(a int, b int) int {
	if b > a || b == a {
		return b
	} else {
		return a
	}
}

func LoopString(String *string) *string {

	temp := strings.Split(*String, "\n")

	tokens := make(map[int]tokenizer.Token)

	for LOOP := 1; LOOP <= len(temp); LOOP++ {
		temp_index := LOOP - 1

		if temp_index == len(temp) {
			break
		}

		tokens = iterString(temp[temp_index], tokens)

		/*
			for _, element := range temp[temp_index] {

				fmt.Printf("TEST -> Line: %d | Element: %c\n", temp_index+1, element)
			}
		*/
	}

	for _, element := range tokens {
		fmt.Println("TYPE -> " + element.GetType().ToString() + " VALUE -> " + element.GetValue())
	}

	return String
}

func iterString(String string, tokens map[int]tokenizer.Token) map[int]tokenizer.Token {

	position := 0
	char := ""
	word := ""

	nextChar := func() {
		char = String[max(position-1, 0):position]
		if char != " " {
			word += char
		}
		position++
	}

	getToken := func() (bool, tokenizer.Token) {
		switch word {
		case "int":
		case "string":
		case "list":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(1), word)
		}

		return true, *tokenizer.CreateToken(0, word)
	}

	for true {

		nextChar()

		fmt.Println(word)

		if char == " " {
			fmt.Println(word)

			created, token := getToken()

			if created {
				tokens[len(tokens)] = token
			}

			word = ""

			continue
		}

		if position > len(String) {
			break
		}

		//mt.Println(word)

		created, token := getToken()

		if created {
			tokens[len(tokens)+1] = token

			continue
		}

	}

	return tokens
}
