package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	//Format specifier
	name := "Alice"
	age := 30
	height := 1.75

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

	num := 42
	fmt.Printf("Decimal: %d\nBinary: %b\nOctal: %o\nHexadecimal: %x\n", num, num, num, num)

	p := Person{"Alice", 30}
	fmt.Printf("%v\n", p) // Prints: {Alice 30}
	num1 := 42
	fmt.Printf("%v\n", num1) // Prints: 42
	fmt.Printf("%+v\n", p)   // Prints: {Name:Alice Age:30}

	//printf and sprintf
	fmt.Printf("Name: %s, Age: %d\n", name, age) // Prints: Name: Alice, Age: 30

	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age) //sprintf is return the formatted string and its not directly print on the console
	fmt.Println(formattedString)                                   // Prints: Name: Bob, Age: 25

	//functions
	//A Simple Function with Parameters and a Return Value
	result := add(5, 3)
	fmt.Println("the result is: ", result)

	//A Simple function that takes a parameter with multiple return a values
	result1, result2 := sumanddifference(5, 4)
	fmt.Println("the sum and difference is: ", result1, result2)

	//A Simple function that takes a no parameter but returns a value
	randomVal := getrandomnum()
	fmt.Println("the random value is: ", randomVal)

	//Anonymous function
	//defining and calling anonymous function
	func() {
		fmt.Println("This is anonymous function")
	}()

	//closure: assigning function to variable and then calling it
	myFunc := func(x, y int) int {
		return x + y
	}
	output := myFunc(4, 5)
	fmt.Println("the addition of two numbers: ", output)

	number := 7

	// Using an anonymous function as a callback
	processCallback(number, func(n int) {
		fmt.Println("Doubled number:", n*2)
	})

}

func add(a int, b int) int {
	return a + b
}

func sumanddifference(a int, b int) (int, int) {
	return a + b, a - b
}

func getrandomnum() int {
	return 42
}

// processCallback takes an integer and a callback function that operates on an integer.
func processCallback(num int, callback func(int)) {
	callback(num)
}
