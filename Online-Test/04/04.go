package main

import "fmt"

var evenCount, oddCount, multiply int

func countEvenNums(numbers []int) int {
	evenCount = 0
	fmt.Print("\nEven Numbers \t: ")
	for _, a := range numbers {
		if a%2 == 0 {
			fmt.Print(a, " ")
			evenCount++
		}
	}
	return evenCount
}
func countOddNums(numbers []int) int {
	oddCount = 0
	fmt.Print("\nOdd Numbers \t: ")
	for _, a := range numbers {
		if a%2 != 0 {
			fmt.Print(a, " ")
			oddCount++
		}
	}
	return oddCount
}

func countMultiplyby5(numbers []int) int {
	multiply = 0
	fmt.Print("\nMultiply By 5 \t: ")
	for _, a := range numbers {
		if a%5 == 0 {
			fmt.Print(a, " ")
			multiply++
		}
	}
	return multiply
}

func main() {
	var size int

	fmt.Print("Enter the size of number (How many?) \t: ")
	fmt.Scan(&size)

	numbers := make([]int, size)

	fmt.Print("Enter number items (From 0 - 1000) \t: ")
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	evenCount = countEvenNums(numbers)
	oddCount = countOddNums(numbers)
	multiply = countMultiplyby5(numbers)
	fmt.Println("\nTotal of Even Numbers \t: ", evenCount)
	fmt.Println("Total of Odd Numbers \t: ", oddCount)
	fmt.Println("Total of Multiply by 5 \t: ", multiply)
}
