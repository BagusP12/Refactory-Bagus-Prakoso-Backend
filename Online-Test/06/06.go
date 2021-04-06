package main

import (
	"fmt"
	"regexp"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func main() {
	fmt.Println(IsLetter("abcdefgh"))
}
