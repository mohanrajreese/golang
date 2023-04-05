package main

import (
	"fmt"
)

func main() {
	// variable declarations

	// without mentioning type
	var myName = "Mohanraj M"
	fmt.Println("My name is ", myName)

	// with mentioning type
	var fatherName string
	fatherName = "Munusamy P"
	fmt.Println(fatherName, " is Father of ", myName)

	// Assigning multiple variable at same time
	var myAge, fatherAge = 24, 48
	fmt.Println("Age of ", myName, " is ", myAge)
	fmt.Println("Age of ", fatherName, " is ", fatherAge)

	var motherName, sisterName string
	motherName = "mariyammal"
	sisterName = "Kiruthika"
	fmt.Println(motherName, ",", sisterName, " are Mother and Sister of ", myName)

	// boolean check
	var isAlive bool
	isAlive = true
	fmt.Println(myName, " is alive : ", isAlive)

	//	Declaring variable without using var
	qualification := "Post Graduate"
	fmt.Println("Qualification of ", myName, " is ", qualification)
}
