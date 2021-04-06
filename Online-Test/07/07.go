package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{9, 4, 2, 4, 1, 5, 3, 0}
	sort.Ints(s)
	fmt.Println(s)
}
