package main

import (
	"fmt"
)

func main() {
	// to create an empty map with a non-zero size,
	// the make syntax is also used with the following:
	// map[KEY_TYPE]VAL_TYPE, i.e. make(map[byte]int)

	newTestMap := make(map[string]string)
	fmt.Println(newTestMap)

	// assigning key value pairs is like in regular language

	newTestMap["key1"] = "value1"
	newTestMap["key2"] = "value2"
	fmt.Println(newTestMap)

	// to get a key-value pair, use the standard map_name["key"]
	findVal := newTestMap["key1"]
	fmt.Println(findVal)

	// find length of a map type data structure
	fmt.Println("the length of the map is:", len(newTestMap))

	// delete a key-value pair from a map
	afterDeleteMap := newTestMap
	delete(afterDeleteMap, "key1")
	fmt.Println("new map is", afterDeleteMap)

	// to check whether a key exists or not
	// use the _ (blank) identifer to return
	// whether a value exists given a key
	_, doesExist := afterDeleteMap["key2"]
	_, doesNotExist := afterDeleteMap["key1"]
	fmt.Println("key2 does exist", doesExist)
	fmt.Println("key1 does not exist", doesNotExist)

	oneLineMapInitalise := map[string]int{"key1": 1, "key2": 3}
	furtherPractice := map[string]byte{"hello-world": 1, "foo": 3}
	fmt.Println("new initalised map", oneLineMapInitalise)
	fmt.Println("further practice", furtherPractice)
}
