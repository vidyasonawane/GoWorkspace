package main

import "fmt"

type myowntype int

var x myowntype = 20

func main() {
	// type myowntype int
	// var x myowntype = 20

	var y int = 25

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//x = y  //not allowed

	x=myowntype(y) //type conversion not casting
	fmt.Println(y)

	

}
