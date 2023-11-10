package main

import "fmt"

func main() {
	// Integer types
	var i int = 42
	fmt.Println("Integer:", i)
	fmt.Printf("i is of type %T", i)
	fmt.Println()

	var j int8 = 12
	fmt.Println("Integer:", j)
	fmt.Printf("j is of type %T", j)
	fmt.Println()

	// Floating-point types
	var f float64 = 3.14
	fmt.Println("Float:", f)
	fmt.Printf("f is of type %T", f)
	fmt.Println()

	// Boolean type
	var b bool = true
	fmt.Println("Boolean:", b)
	fmt.Printf("b is of type %T", b)
	fmt.Println()

	// String type
	var s string = "Hello, Go!"
	fmt.Println("String:", s)
	fmt.Printf("s is of type %T", s)
	fmt.Println()

	// Character type (represented by runes)-alias int32
	var r rune = 'A'
	fmt.Println("Rune:", r)
	fmt.Printf("r is of type %T", r)
	fmt.Println()

	// Complex number type
	var c complex128 = complex(1, 2)
	fmt.Println("Complex:", c)
	fmt.Printf("c is of type %T", c)
	fmt.Println()

	// Arrays
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slices
	slice := []int{4, 5, 6}
	fmt.Println("Slice:", slice)

	// Maps
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map:", m)

	// Structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Alice", 30}
	fmt.Println("Struct:", person)
}
