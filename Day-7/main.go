package main

import "fmt"

func main() {
	//Identifiers
	// Valid variable names
	var age int
	var user_name string
	var totalAmount float64

	// Assign values to variables
	age = 25
	user_name = "John Doe"
	totalAmount = 100.50

	// Print variable values
	fmt.Println("Age:", age)
	fmt.Println("User Name:", user_name)
	fmt.Println("Total Amount:", totalAmount)

	// Invalid variable names
	//var 123abc int       // starts with a digit
	//var user-name string // contains a hyphen
	//var totalAmount$ float64 // contains a special character
	//var func int   //contains a keyword func

	//Operators
	// Arithmetic operators
	a := 10
	b := 4
	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division
	fmt.Println("a % b =", a%b) // Modulus

	// Comparision operators
	fmt.Println("a == b:", a == b) // Equal to
	fmt.Println("a != b:", a != b) // Not equal to
	fmt.Println("a > b:", a > b)   // Greater than
	fmt.Println("a < b:", a < b)   // Less than
	fmt.Println("a >= b:", a >= b) // Greater than or equal to
	fmt.Println("a <= b:", a <= b) // Less than or equal to

	// Logical operators
	x := true
	y := false
	fmt.Println("x && y:", x && y) // Logical AND
	fmt.Println("x || y:", x || y) // Logical OR
	fmt.Println("!x:", !x)         // Logical NOT

	// Assignment operators
	c := 5
	c += 3 // Equivalent to c = c + 3
	fmt.Println("c =", c)

	//Bitwise Operator
	// Bitwise AND
	resultAND := 5 & 3
	fmt.Println("5 & 3 = ", resultAND)
	// Bitwise OR
	resultOR := 5 | 3
	fmt.Println("5 | 3 = ", resultOR)
	// Bitwise XOR
	resultXOR := 5 ^ 3
	fmt.Println("5 ^ 3 = ", resultXOR)

	// Unary Operator
	d := 8
	d++ // Increment by 1
	fmt.Println("d after increment:", d)
	d-- // Decrement by 1
	fmt.Println("d after decrement:", d)

	//For loop
	// Basic for loop: The loop initializes i to 0, executes the loop as long as i is less than 5, increments i in each iteration, and prints the value of i
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop used as a while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite for loop
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break // exit the loop when k is 5
		}
	}

	// For loop over a collection : we use the range keyword to iterate over elements in a slice (numbers). The loop returns both the index and value of each element in the slice.
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}
