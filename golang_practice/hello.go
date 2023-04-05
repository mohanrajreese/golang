package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

func multiply(a, b int) int {
	return a * b
}

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println("Hello Mohanraj")
	fmt.Println("current ime is, ", time.Now())
	fmt.Println("my favorite number is,", rand.Intn(10))
	fmt.Printf("Now we have more than %g problems.\n", math.Sqrt(9))
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(math.Pi))

	// Add function used here
	fmt.Println("Add two numbers: ", add(12, 24))

	// Multiply function used here
	fmt.Println("Multiply two numbers:", multiply(42, 2))

	// Swap function used here
	a, b := swap("hai", "mohan")
	fmt.Println(a, b)

}
