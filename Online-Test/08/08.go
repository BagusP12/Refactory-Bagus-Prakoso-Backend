package main

import (
	"fmt"
	"strings"
)

func main() {
	value := "23dn3ir30fd2eddd"
	fmt.Println(value)

	result := strings.Replace(value, "23dn3ir30fd2e", "*************", -1)
	fmt.Println(result)
}
