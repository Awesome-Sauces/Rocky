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

	position := 0
	char := ""
	word := ""

	nextChar := func() {
		char = String[max(position-1, 0):min(position, len(String))]
		if char != " " {
			word += char
		}
		position++
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

		default:
			return false
		}
	}

	getToken := func(words string) (bool, tokenizer.Token) {
		words = strings.TrimSpace(words)

		switch words {
		case "int":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(11), words, 0)
		case "string":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(11), words, 0)
		case "double":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(11), words, 0)
		case "list":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(11), words, 0)
		case "+":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(1), words, 0)
		case "-":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(2), words, 0)
		case "*":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(4), words, 0)
		case "/":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(3), words, 0)
		case "=":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(15), words, 0)
		case "(":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(5), words, 0)
		case ")":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(6), words, 0)
		case "{":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(8), words, 0)
		case "}":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(7), words, 0)
		case "#":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(12), words, 0)
		case "[":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(14), words, 0)
		case "]":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(13), words, 0)
		case ";":
			return true, *tokenizer.CreateToken(tokenizer.TOKEN(16), words, 0)
		}

		if len(words) == 0 {
			return false, *tokenizer.CreateToken(0, words, 0)
		}

		return true, *tokenizer.CreateToken(10, words, 0)

		//return false, *tokenizer.CreateToken(0, words, 0)
	}

	for true {

		nextChar()

		var created bool
		var token tokenizer.Token

		if char == " " ||
			isSingleton(char) {
			word = strings.ReplaceAll(word, " ", "")
			word = strings.ReplaceAll(word, char, "")

			created, token := getToken(word)

			if created {
				tokens[len(tokens)] = &token
			}

			created2, token2 := getToken(char)

			if created2 {
				tokens[len(tokens)] = &token2
			}

			word = ""

			continue
		}

		if position-1 > len(String) {
			break
		}

		if isSingleton(char) {
			word = strings.ReplaceAll(word, char, "")
			created, token = getToken(char)
		}

		if created {

			if created {
				tokens[len(tokens)] = &token
			}

			word = ""

			continue
		}

	}

	return tokens
}
