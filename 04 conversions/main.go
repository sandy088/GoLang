package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter ratings")
	ratings := bufio.NewReader(os.Stdin)
	rating, _ := ratings.ReadString('\n')

	numRating, err := strconv.ParseInt(strings.TrimSpace(rating), 10, 64)

	if err != nil {
		fmt.Println("Error parsing rating:", err)
		return
	} else {
		fmt.Println("Here are the ratings ", numRating)
		fmt.Printf("Type of numRating is %T\n", numRating)
	}

}
