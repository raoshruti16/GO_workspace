package main

import (
	"fmt"
	"strconv"
)

// BubbleSort use bubble sort to sort given slice
func BubbleSort(integerSlice []int) {
	length := len(integerSlice)
	for i := 0; i < length-1; i++ {
		modified := false
		for j := 0; j < length-i-1; j++ {
			if integerSlice[j] > integerSlice[j+1] {
				Swap(integerSlice, j)
				modified = true
			}
		}

		if !modified {
			break
		}
	}
}

// Swap to exchange two values at position i of given slice
func Swap(integerSlice []int, i int) {
	integerSlice[i], integerSlice[i+1] = integerSlice[i+1], integerSlice[i]
}

func main() {
	var integerSlice = make([]int, 0, 10)
	var s string

	for i := 0; i < 10; i++ {
		fmt.Println("Enter 10 integers to sort")

		fmt.Scan(&s)

		num, _ := strconv.Atoi(s)
		integerSlice = append(integerSlice, num)

		BubbleSort(integerSlice)

		fmt.Println(integerSlice)

	}
}
