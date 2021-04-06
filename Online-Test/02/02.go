package main

import (
	"fmt"
	"strings"
)

func CensorWord(str string, censored []string) string {

	var newSlice []string

	if len(censored) <= 0 {
		return str
	}

	strSlice := strings.Fields(str)

	for position, word := range strSlice {
		for _, forbiddenWord := range censored {

			if test := strings.Index(strings.ToLower(word), forbiddenWord); test > -1 {
				//if test := strings.EqualFold(strings.ToLower(word), forbiddenWord); test == true {

				replacement := strings.Repeat("*", len(word))

				strSlice[position] = replacement
				newSlice = append(strSlice[:position], strSlice[position:]...)
			}
		}
	}

	return strings.Join(newSlice, " ")

}

func main() {
	var censoredwords = []string{"dolor",
		"elit",
		"quis",
		"nisi",
		"fugiat",
		"proident",
		"laborum"}

	var inputString = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	resultString := CensorWord(inputString, censoredwords)

	fmt.Println("[Original : \n]", inputString)
	fmt.Println("-------------------------------------------")
	fmt.Println("[Censored : \n]", resultString)

}
