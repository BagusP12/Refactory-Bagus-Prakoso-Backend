package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 12, 4, 5, 8, 9}
	sort.Ints(s)
	fmt.Println(s)
}
