package main

import (
	. "fmt"
	. "time"
)

func main() {
	var number int
	Println("Enter the input :")
	var input, err = Scanln(&number)
	if err != nil {
		println("There is an error in the input")
		return
	}
	println(input, "A")
	if number%2 == 0 {
		Println("The given input is Even Number", number)
	} else {
		Println("The given number is Odd number : ", number)
	}

	if number < 0 {
		Println("the given number is negative number ", number)
	} else if number == 0 {
		Println("the give number is Zero")
	} else {
		Println("the given number is positive number", number)
	}

	Println("The Output is: ")
	Println(number)

	Println(Now().Weekday())
}
