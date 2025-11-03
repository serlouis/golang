package main

import (
	"fmt"
	"maps"
)

func main(){

	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)
	// using a map literal
	//mapVariable = map[keyType]valueType{
	//key1: value1,
	//key2: value2
	//}

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"]) // wrong key so will return zero
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	//clear(myMap)
	//fmt.Println(myMap)

	_, ok := myMap["key1"]
	//fmt.Println(value) // returns 9 as expected
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	//fmt.Println("Is a value associated with key1 ", ok) // returns true this means there is a value associated with this key in the map this is useful for making apis just to check if there si a value that is associated with the key

	// initialize and declare a new map in the same line
	myMap2 := map[string]int{"a":1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a":1, "b":2}

	// maps also have an equality check

	if maps.Equal(myMap3, myMap2){
		fmt.Println("myMap3 and myMap2 are equal")
	}

	// iterating over a map
	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	// this will iterate over the values and doesn't print the key this may be useful
	for _, v := range myMap3 {
		fmt.Println(v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initalized to nil value.")
	}

	// try to read from nil map
	val := myMap4["key"]
	fmt.Println(val) // should return blank string

	// if the map is initialized to nil have to use the make function to initialize
	// write to nil map
	//myMap4["key"] = "Value" // this won't work
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value" // now this works
	fmt.Println(myMap4)

	// we can get the length of a map using the len function the length is the number of keys in the map not the number of values
	fmt.Println("myMap4 length is" , len(myMap))

	// nested maps an outer map can have inner maps embedded inside it
	myMap5 := make(map[string]map[string]string)
	
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
	