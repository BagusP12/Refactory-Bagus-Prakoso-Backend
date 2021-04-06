package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "aku ingin begini " +
		"aku ingin begitu " +
		"ingin ini itu banyak sekali " +
		"semua semua semua " +
		"dapat dikabulkan " +
		"dapat dikabulkan " +
		"dengan kantong ajaib " +
		"aku ingin terbang bebas " +
		"di angkasa " +
		"hei… baling baling bambu " +
		"la... la... la... " +
		"aku sayang sekali " +
		"doraemon " +
		"la... la... la... " +
		"aku sayang sekali "

	fmt.Println("Original : \n", str)

	res1 := strings.Count(str, "aku")
	res2 := strings.Count(str, "ingin")
	res3 := strings.Count(str, "dapat")

	fmt.Println("\nAku	: ", res1)
	fmt.Println("Ingin	: ", res2)
	fmt.Println("Dapat	: ", res3)

}
