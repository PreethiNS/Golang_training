package main

import "fmt"

func main() {
	//Example: Using break to exit a loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break // exit the loop when i is 5
		}
	}

	// Example: Using continue to skip an iteration
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // skip the iteration when i is 2
		}
		fmt.Println(i)
	}

	//Array
	// Declaring and initializing an array and the size is fixed
	var numbers [5]int // An array of integers with a length of 5 and intialize to default values :the default value if integer is 0

	// Assigning values to array elements
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	//number[5] = 6 - index out of range because the size of an array is 5

	// Accessing and printing array elements
	fmt.Println("Element 0:", numbers[0])
	fmt.Println("Element 1:", numbers[1])

	//Decalre and intialize in same line
	a := [4]int{1, 2, 3, 4}
	fmt.Println("the first element in array a is : ", a[0])
	//length of an array
	fmt.Println("length of array a is: ", len(a))

	//Array act as a Pass by value
	//assigning one array to another: original array a to b array
	b := a
	fmt.Println("array b is: ", b)
	//Modifying an element in b doesn't affect the original array a
	b[0] = 100
	fmt.Println("array b after changes in 0th index is: ", b)
	fmt.Println("array a after done changes in array a(0th index): ", a)

	//Slice
	// Creating a slice
	var n []int // This creates a nil slice

	// Appending elements to the slice
	n = append(n, 1)
	n = append(n, 2, 3, 4)

	fmt.Println("slice n is: ", n)
	// Accessing and printing slice elements
	fmt.Println("Element 0:", n[0])
	fmt.Println("Element 1:", n[1])
	fmt.Println("length and capacity of slice n is: ", len(n), cap(n))

	// Initializing a slice with specific values
	m := []int{1, 2, 3, 4}

	// Printing slice elements
	fmt.Println("Slice elements:")
	for i, num := range m {
		fmt.Printf("Element %d: %d\n", i, num)
	}

	// Creating a slice with an initial capacity
	num := make([]int, 3, 5) // The slice has a length of 3 and a capacity of 5
	fmt.Println("length and capacity of slice n is: ", len(num), cap(num))
	num[0] = 9
	num[1] = 8
	num = append(num, 1, 2, 3, 4)
	fmt.Println("the slice num is: ", num)
	//dynamically the length and capacity gets changed- when you append more elements than the capacity defined before
	fmt.Println("length and capacity of slice n is: ", len(num), cap(num))

	// Creating a slice
	slice1 := []int{1, 2, 3, 4, 5}

	// Assigning one slice to another
	slice2 := slice1

	// Modifying an element in slice2 affect the slice1 and vice versa
	slice2[0] = 10

	// Printing both slices
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

}
