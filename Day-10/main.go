package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Id     int
	Person Person //Embedded struct - name and type/struct_name
	Salary int
}

type Employee1 struct {
	Id     int
	Person //Embedded struct - if we specified type/struct_name directly without explicitly specify its name as a field then we can access Name value using Employee1 instance i.e e1.Name instead of e1.Person.Name
	Salary int
}

func main() {
	//1.creating an instance of type Person struct
	p := Person{
		Name: "xxx",
		Age:  23,
	}
	fmt.Println("p: ", p)                //print only values
	fmt.Printf("p : %+v\n", p)           //print both key:value's
	fmt.Println("The name is: ", p.Name) //Accessing the value to the field Name of the struct

	//2. another way of creating an instance of type Person struct
	var p1 Person

	//assigning values to the fields of struct Person
	p1.Name = "zzz"
	//if you won't assign value to struct field then it holds its default value ex: Age which is of type integer - default value 0 as I'm not initializing any value to Age field
	fmt.Println("p1: ", p1)

	//after assigning value to Age field
	p1.Age = 24
	fmt.Println("after assignment to age p1: ", p1)
	fmt.Printf("after assignment to age p1: %+v\n", p1) //print key:value

	//struct tags - we will understand in seperate session
	/*
			type Person struct {
				Name string  `json:"name"`
				Age  int	 `json:"age"`
		}
	*/

	//Embedded struct
	e := Employee{
		Id: 1,
		Person: Person{
			Name: "rrr",
			Age:  26,
		},
		Salary: 35000,
	}
	fmt.Printf("e: %+v\n", e)
	//access Name value from Person struct using Employee instance that is e
	fmt.Println("The name is: ", e.Person.Name)
	//access Id value which belongs to Employee struct using the Employee instance that is e
	fmt.Println("The Id is: ", e.Id)

	//creating an instance of Employee1 struct
	e1 := Employee1{
		Id: 2,
		Person: Person{
			Name: "qqq",
			Age:  27,
		},
		Salary: 40000,
	}
	fmt.Printf("e1: %+v\n", e1)
	fmt.Printf("The name is: %+v\n", e1.Name)

	//Comments
	//This is single line comment - non executable

	/*
		This is
		Multi line comment - non executable
	*/

}
