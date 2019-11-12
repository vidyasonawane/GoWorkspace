package main

import "fmt"

func main() {
	mymap := map[string]int{
		"vidya":  22,
		"kunal":  19,
		"Prasad": 27,
		"Sayali": 30,
	}
	fmt.Println(mymap)

	//range to print all key value pairs
	for k, v := range mymap {
		fmt.Println(k, v)
	}

	//Print value for a given key
	fmt.Println(mymap["vidya"])
	fmt.Println(mymap["kunal"])
	fmt.Println(mymap["baba"]) //It will print zero value as it is not in map

	//To check wether the value exists in map or not.
	val, status := mymap["baba"]
	fmt.Println(val)    //assigned value if exists, zero value if not
	fmt.Println(status) //true if exists, false if not

	//If statement to check the status of key
	if val, status := mymap["vidya"]; status {
		fmt.Println("Key exists with value", val)
	}

	//adding the element
	mymap["Pihu"] = 2

	fmt.Println(mymap)

	//deleting an element
	delete(mymap, "vidya")
	fmt.Println(mymap)
	delete(mymap, "pihu") //case sencetive
	fmt.Println(mymap)
	delete(mymap, "Pihu") //case sencetive
	fmt.Println(mymap)

}
