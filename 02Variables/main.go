package main

import "fmt"

//⚠️ I have keep the L capital of Login Token to make it public -> In go lang if any variable has capital letter in the beginning then it is public
const LoginToekn string = "1234567890"

func main() {
	//Go automatically assigns default values to variables to 0 -> That's the beauty of GOlang
	var myVariable int
	fmt.Println(myVariable)
	fmt.Printf("Type of myVariable is %T\n", myVariable)

	//Most important syntax
	mysecVariable := "10"
	fmt.Println(mysecVariable)
	fmt.Printf("Type of myVariable is %T\n", mysecVariable)
}
