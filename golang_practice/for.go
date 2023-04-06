package main

import "fmt"

func main() {
	// Basic loop
	fmt.Println("bassic for loop")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i += 2
	}

	// classic for loop
	fmt.Println("classic for loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		for k := 1; k <= 5; k++ {
			fmt.Println(k, " time is 12.30 AM")
		}
		break
	}

	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			fmt.Println("Even number", n)
			continue
		}
		fmt.Println("Odd number", n)
	}
}
