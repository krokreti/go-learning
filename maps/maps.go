package main

import "fmt"

func main()  {
	//to create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)

	//set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	//printing a map with s.g. fmt.Println will show all of its key/value pairs
	fmt.Println("map:" , m)

	//get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if the key doesnt exist, the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// the builtin len returns the number of key/value pairs from a map
	fmt.Println("len:", len(m))

	//the builtin delete removes key/values pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	/*
	The optional second return value when getting a value from a map indicates if the key was present in the map. 
	This can be used to disambiguate between missing keys and keys with zero values like 0 or "". 
	Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	*/
	_, prs := m["k2"]
    fmt.Println("prs:", prs)

	//You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}