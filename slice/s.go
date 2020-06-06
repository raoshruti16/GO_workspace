package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var slice []int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(`Enter integer: `)
	for scanner.Scan() {
		input := scanner.Text()
		if input == `X`{
			break
		}
		intNum,_ := strconv.Atoi(input)
		slice = append(slice,intNum)
		fmt.Print(`Enter integer: `)
	}

	 sort.Ints(slice)

	fmt.Println(slice)
}
