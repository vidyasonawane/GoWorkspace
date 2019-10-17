package main

import "fmt"

func main(){
	a := "stored in a"
	//b := "stored in b"
	fmt.Println("a = ",a)
	//b is not used - invalid code and will give an error: b declared and not used
	
	n, _  := fmt.Println("Vidya") //blank identifier is used to avoid printing an error.
	fmt.Println(n)
}