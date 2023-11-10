package main

import "fmt"

func main() {
	//Explicit conversion
	var x float64 = 3.14
	//converting float64 to int
	y := int(x)
	fmt.Println("y value is :", y)

	//Control flow
	//Example 1: if statement
	z := 7

	if z > 5 {
		fmt.Println("z is greater than 5")
	}

	//Example 2: if-else statement
	a := 3

	if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is not greater than 5")
	}

	//Example 3: if-else if statement
	b := 12

	if b > 15 {
		fmt.Println("b is greater than 15")
	} else if b > 10 {
		fmt.Println("b is greater than 10 but not greater than 15")
	} else {
		fmt.Println("b is 10 or less")
	}

	// Example: Initialize and condition in the same if block
	if num := 9; num > 5 {
		fmt.Println("num is greater than 5")
	} else {
		fmt.Println("num is not greater than 5")
	}

	//Switch statement
	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's a weekday.")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("Invalid day.")
	}
}
