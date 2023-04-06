package main

import (
	"fmt"
	"time"
)

func main() {
	var number int
	fmt.Println("Enter the input :")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("The given input is Even Number", number)
	} else {
		fmt.Println("The given number is Odd number : ", number)
	}

	if number < 0 {
		fmt.Println("the given number is negative number ", number)
	} else if number == 0 {
		fmt.Println("the give number is Zero")
	} else {
		fmt.Println("the given number is positive number", number)
	}

	fmt.Println("The Output is: ")
	fmt.Println(number)

	fmt.Println(time.Now().Weekday())
}
