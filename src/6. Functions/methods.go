package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct{
	person
	ltk bool
}



func main(){
	sa1 := secretAgent {
		person:person{
			first:"vidya",
			last:"sonawane",
		},
		ltk:true,
	}
	sa2 := secretAgent {
		person:person{
			first:"kunal",
			last:"patil",
		},
		ltk:false,
	}

	sa1.speak()
	sa2.speak()

}

func (r secretAgent) speak(){
	fmt.Println("Hey",r.first,r.last)
}