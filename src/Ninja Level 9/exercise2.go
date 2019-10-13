package main
import "fmt"
//method sets
type person struct {
	first string
	age int
}

type human interface {
	speak()
}

func (p *person) speak(){
	fmt.Println("Hello", p.first)
}

func saySomething(h human){
	h.speak()
}

func main(){
	p1 := person{
		first : "vidya",
		age: 22,
	}
	saySomething(&p1)
	//saySomething(p1) //not work
	//p1.speak() //will work
}