package main

import (
	"fmt"

	"github.com/PreethiNS/Variables/package1"
)

// 2.Package level variable
var package_var int = 30

func main() {
	//1. Local Variable
	var local_var int = 20
	fmt.Println("local variable is : ", local_var)

	//print the package level variable
	fmt.Println(package_var)

	//Print the exported variable which declared and initialized in the package1
	fmt.Println(package1.Exported_var)

	//Constants declaration and initialzation
	const pi = 3.14
	fmt.Println("pi value is: ", pi)
	/*
		you can't change constant value i.e pi- result an error
		pi=3.14567
	*/

	var x int = 10
	fmt.Println("x: ", x)   //10
	fmt.Println("&x: ", &x) //print memory address of x

	//Pointers: declaration and intialization
	var ptr *int = &x
	fmt.Println("pointer address: ", &ptr)        //print memory address pointer "ptr"
	fmt.Println("x address using pointer: ", ptr) //print x address indirectly using pointer
	fmt.Println("x value using pointer: ", *ptr)  //print x value indirectly using pointer

	var ptr1 *int = &x
	fmt.Println("pointer address: ", ptr1)
	//arithmetic operation cannot be performed on pointers ex: ptr+ptr1 - result an error
	//you can perform addition on values not on address
	fmt.Println("addition of values accessed by pointers: ", *ptr+*ptr1)

	passvariable(x)
	//x value doesn't change because num1 is copy of the x; the changes made to num1 doesn't affect the x
	fmt.Println("after passvariable x value is: ", x)

	passpointer(&x)
	//x value changes because the function worked with the original data(the value present in x)through the pointer
	fmt.Println("after passpointer x value is: ", x)

	//Nil pointer
	var pt *int
	//without intialization - default value of pointer is nil
	fmt.Println("pt default value is: ", pt)

}
func passvariable(num1 int) {
	num1 = num1 * 2
	fmt.Println("num1: ", num1)
}

func passpointer(num1 *int) {
	*num1 = *num1 * 2
	fmt.Println("num1: ", *num1)
}
