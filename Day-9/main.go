package main

import "fmt"

func main() {
	//1.map decalaration using make function
	myMap := make(map[string]int)

	//Adding and accessing elements
	myMap["apple"] = 10
	myMap["banana"] = 20

	fmt.Println("number of apples: ", myMap["apple"])

	//deleting elements
	delete(myMap, "apple")

	//checking for existence
	value, exist := myMap["apple"]
	if exist {
		fmt.Println("the value of apples: ", value)
	} else {
		fmt.Println("value not found")
	}

	//Itertaing over map
	for key, value := range myMap {
		fmt.Printf("key: %s value: %d", key, value)
	}
	fmt.Println()

	//map as reference type
	m1 := myMap
	fmt.Println("The map m1 is : ", m1)
	m1["Kiwi"] = 30
	fmt.Println("after add kiwi to m1: ", m1)
	fmt.Println("the changes made to myMap: ", myMap)

	//2.Map declaration using literal
	anotherMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("anotherMap is : ", anotherMap)
	fmt.Println("The first key value is : ", anotherMap["key1"])

}
