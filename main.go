package main

import "fmt"

func main() {
	//1.using var
	//declaration
	var x int
	fmt.Println(&x)
	fmt.Println(x)
	//assignment or initialization
	x = 10
	fmt.Println(x)

	//declare and initialize in the same line
	var y int = 20
	fmt.Println(y)

	//2. Type inference with Var
	var z = 30
	fmt.Println(z)

	//3. short variable declaration
	a := 40
	fmt.Println(a)

}
