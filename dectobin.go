package main

import "fmt"
import "strconv"

func main() {
	var dec int64
	fmt.Print("Enter your Decimal Value: ")
	fmt.Scanln(&dec)
	output := strconv.FormatInt(dec, 2)
	fmt.Print("Binary Value is ", output)
	
}