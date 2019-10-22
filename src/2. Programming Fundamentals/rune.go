package main

import "fmt"

func main(){
	foo := 'a'
	fmt.Println(foo) //it will print the utf8 of a
	fmt.Printf("%T \n",foo)
	fmt.Printf("%T \n",rune(foo))
}