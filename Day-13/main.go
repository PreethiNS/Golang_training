package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type myInt int

func main() {
	//Create a instance of Person type
	person := Person{"Alice", 20}
	//Call sayHello method on the Person Instance(i.e person)
	person.sayhello()

	m := myInt(10)
	out := m.Double()
	fmt.Println("the output is: ", out)

	//Reciever type
	p1 := Person{"alice", 20}
	//Value reciever
	p1.sayhello()
	fmt.Println("after changing name: ", p1.Name) //value doesnt change

	//Pointer reciever
	p1.changeName()
	fmt.Println("after changing name: ", p1.Name) //value changes

	//Error handling
	val, err := divide(4, 0)
	if err != nil {
		fmt.Println("the error is: ", err)
	} else {
		fmt.Println("division result: ", val)
	}

}

// Method:value reciever
func (p Person) sayhello() {
	fmt.Printf("hello my name is: %s and I am %d years old.\n", p.Name, p.Age)
	p.Name = "yyy"
	fmt.Printf("hello my name is: %s and I am %d years old.\n", p.Name, p.Age)

}

func (val myInt) Double() myInt {
	return val * 2
}

//Method : Pointer reciever
func (p *Person) changeName() {
	p.Name = "xxx"
	fmt.Printf("hello my name is: %s and I am %d years old.\n", p.Name, p.Age)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
