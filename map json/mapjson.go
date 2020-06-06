package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// init a map
	addressMap := make(map[string]string)
	// create new reader
	reader := bufio.NewReader(os.Stdin)

	// get user input for name
	fmt.Printf("Enter a name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	// get user input for address
	fmt.Printf("Enter an address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSuffix(address, "\n")

	// add name and address to map
	addressMap["name"] = name
	addressMap["address"] = address

	// convert map to json by Marshal, and print json string
	jsonObject, _ := json.Marshal(addressMap)
	fmt.Println(string(jsonObject))

}
