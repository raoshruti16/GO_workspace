package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	size := 0
	sortedArray := make([]int, 3)

	for {
		fmt.Scan(&input)
		if input == "X" {
			break;
		}
		number, err:= strconv.Atoi(input)
		if err == nil {
			sortedArray = insertToArray(number, sortedArray, size)
			size = size + 1
			printAllValues(sortedArray, size)
		} else {
			println("Invalid Input, enter X to exit")
		}
	}
}
func printAllValues(sortedArray []int, size int) {
	print("Updated Slice:")
	for i := 0; i < size; i++ {
		fmt.Printf(" %d", sortedArray[i])
	}
	println()
}
func insertToArray(input int, sortedArray []int, size int) []int {

	limit := len(sortedArray)
	if size < len(sortedArray) {
		limit = size
	}

	for i := 0; i < limit; i++ {
		if sortedArray[i] > input {
			temp := sortedArray[i]
			sortedArray[i] = input
			input = temp
		}
	}
	if size >= 3 {
		sortedArray = append(sortedArray, input)
	} else {
		sortedArray[size] = input
	}
	return sortedArray
}
