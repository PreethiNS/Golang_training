package main

import "fmt"

func main() {

	//Variadic function
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("the sum is: ", result)

	//Underscore: ignore the value
	result1, _ := sumanddifference(5, 3)
	fmt.Println("the result1 is: ", result1)

	//defer
	deferexample()

	//panic
	/*doSomething()

	fmt.Println("after the panic")*/

	//Comma,ok
	fruits := map[string]int{"Apple": 4, "Banana": 5}

	if counts, exists := fruits["Apple"]; exists {
		fmt.Println("the number of apples: ", counts)
	} else {
		fmt.Println("The apples are not found")
	}

}

func sum(num ...int) int {
	add := 0

	for _, val := range num {
		add += val
	}
	return add
}

func sumanddifference(a int, b int) (int, int) {
	return a + b, a - b
}

func deferexample() {
	defer fmt.Println("this is last statement")

	fmt.Println("first statement")
	fmt.Println("second statement")

}

//uncomment when you want to execute to see the effect of panic
/*func doSomething(){

	//recover when panic occurs
	/*
	defer func(){
		if r:= recover();r!=nil{
			fmt.Println("Recovered: ",r)
		}

	}()
	//simulating panic
	panic("oops something wrong")
}*/
