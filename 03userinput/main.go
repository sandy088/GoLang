package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey what's your name?")
	reader := bufio.NewReader((os.Stdin))

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Hello", input)
}
