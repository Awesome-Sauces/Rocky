package lexer

import (
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

func min(a int, b int) int {
	if b < a || b == a {
		return b
	} else {
		return a
	}
}

func LoopString(String *string) map[int]*tokenizer.Token {

	temp := strings.Split(*String, "\n")

	tokens := make(map[int]*tokenizer.Token)

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

	for index, element := range tokens {
		element.SetOrder(index)
	}

	return tokens
}

func iterString(String string, tokens map[int]*tokenizer.Token) map[int]*tokenizer.Token {

	//maxPosition := len(String)
	position := 0
	char := ""
	word := ""

	/*backward := func() {
		if position > 0 &&
			maxPosition > 0 {
			position -= 2
			maxPosition -= 2

			char = String[max(position-1, 0):min(position, len(String))]

		}
	}*/

	nextChar := func() {
		char = String[max(position-1, 0):min(position, len(String))]
		word += char

		position++
	}

	isDigit := func(character string) bool {
		switch character {
		case "1":
			return true
		case "2":
			return true
		case "3":
			return true
		case "4":
			return true
		case "5":
			return true
		case "6":
			return true
		case "7":
			return true
		case "8":
			return true
		case "9":
			return true
		case "0":
			return true
		default:
			return false
		}
	}

	isSingleton := func(character string) bool {
		switch character {
		case "+":
			return true
		case "-":
			return true
		case "*":
			return true
		case "/":
			return true
		case "=":
			return true
		case "(":
			return true
		case ")":
			return true
		case "{":
			return true
		case "}":
			return true
		case "#":
			return true
		case "[":
			return true
		case "]":
			return true
		case ";":
			return true
		case ".":
			return true

		default:
			return false
		}
	}

	builder := false

	for true {

		nextChar()

		// This number detection code is a little buggy when encountering the builder
		// Being used by a string constructor
		if isDigit(char) && !builder {
			builder = true
		} else if builder && !isDigit(char) && char != "." && !strings.Contains(word, "\"") {
			tokens[len(tokens)] = tokenizer.CreateToken(tokenizer.NUMBER, strings.ReplaceAll(word, char, ""), 0)
			word = char

			builder = false
		}

		// Works perfectly solo but needs
		// More compatibility with Numbers
		if char == "\"" {
			if builder {
				tokens[len(tokens)] = tokenizer.CreateToken(tokenizer.STRING, word, 0)
				word = ""
			} else {
				builder = true
			}
		}

		if isSingleton(word) {
			tokenType := tokenizer.FromString(word)

			if tokenType.ToString() != "NONE" {
				tokens[len(tokens)] = tokenizer.CreateToken(tokenType, word, 0)
			}

			word = ""
		}

		if !builder && (isSingleton(char) || char == " ") && len(word) > 1 {
			word = strings.ReplaceAll(word, char, "")
			tokenType := tokenizer.FromString(word)

			if tokenType.ToString() != "NONE" {
				tokens[len(tokens)] = tokenizer.CreateToken(tokenType, word, 0)
			}

			tokenType = tokenizer.FromString(char)

			if tokenType.ToString() != "NONE" {
				tokens[len(tokens)] = tokenizer.CreateToken(tokenType, char, 0)
			}

			word = ""

			continue
		}

		if position > len(String) {
			break
		}

	}

	return tokens
}
