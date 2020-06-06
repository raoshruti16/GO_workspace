/*
This programs aims to separate a slice in 4 and concurrently sort the content
in it.

Sadly I couldn't find a way to merge the 4 sorted slices without sorting yet
again. So this code is very inefficient that just using sort.Ints.

Yet, it showcase how to concurrently run code and use the result for
another task.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func prompt() []int {
	var numbers []int

	fmt.Println("Please enter the numbers you want to sort, separeted by commas")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	for _, r := range input {
		n, _ := strconv.Atoi(r)
		numbers = append(numbers, n)
	}
	return numbers
}

func order(numbers []int) (sortedNumbers []int) {
	var chunks [][]int

	maxChunks := 4
	chunkSize := (len(numbers) + maxChunks - 1) / maxChunks

	for i := 0; i < len(numbers); i += chunkSize {
		j := i + chunkSize
		if j > len(numbers) {
			j = len(numbers)
		}
		chunks = append(chunks, numbers[i:j])
	}

	fmt.Println("Input separated in 4 chunks: ", chunks)

	var wg sync.WaitGroup
	for _, c := range chunks {
		wg.Add(1)
		go sortChunk(c, &wg)
	}
	wg.Wait()

	for _, c := range chunks {
		sortedNumbers = append(sortedNumbers, c...)
	}

	sort.Ints(sortedNumbers)

	return
}

func sortChunk(chunk []int, wg *sync.WaitGroup) {
	sort.Ints(chunk)
	fmt.Println("Sorted chunk", chunk)
	wg.Done()
}

func main() {
	// numbers := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	numbers := prompt()
	orderedNumbers := order(numbers)
	fmt.Println("List of numbers", orderedNumbers)
}