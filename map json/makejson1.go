package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	print("Enter you name: ")
	name := readStringFromConsole()
	print("Enter you address: ")
	address := readStringFromConsole()

	addressBook := make(map[string]string)
	addressBook[name] = address

	jsonString, _ := json.Marshal(addressBook)
	fmt.Println(string(jsonString))
}

func readStringFromConsole() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = strings.Trim(input, " \n")
	return input
}