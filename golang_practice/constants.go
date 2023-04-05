// Declare the main package.
package main

// Import necessary packages.
import (
	"fmt"
	"math"
)

// Declare a constant variable 'name'.
const name = "Mohanraj"

// Define the main function.
func main() {
	// Print the value of the 'name' constant.
	fmt.Printf(name)

	// Declare and initialize a constant variable 'age'.
	const age = 500000000

	// Print the age.
	fmt.Println("age is ", age)

	// Declare and initialize a constant variable 'devide'.
	const devide = 3e20 / age

	// Print the value of 'devide'.
	fmt.Println("devide = ", devide)

	// Print the value of 'devide' as an int64.
	fmt.Println(int64(devide))

	// Print the sine of 'age'.
	fmt.Println(math.Sin(age))
}
