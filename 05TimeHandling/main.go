package main

import (
	"fmt"
	"time"
)

func main() {

	presenttime := time.Now()

	fmt.Println("Time is ", presenttime)
	fmt.Println("Time is ", presenttime.Format("01-02-2006, 15:04 Monday"))

	t := time.Date(2020, time.January,
		11, 21, 34, 01, 0, time.UTC)

	fmt.Printf("Go launched at %s\n", t.Local())
	fmt.Println("Created Date is ", t)
}
