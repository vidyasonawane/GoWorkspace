package main

import "fmt"

const (
	
	_ = iota
	kb = 1 << (iota*10)
	mb = 1 << (iota*10)
	gb = 1 << (iota*10)

)

func main() {
	a := 45
	fmt.Printf("%d\t%b\n",a,a)
	b := a << 1
	fmt.Printf("%d\t%b\n",b,b)
	c := a << 2
	fmt.Printf("%d\t%b\n",c,c)

	d := a >> 1
	fmt.Printf("%d\t%b\n",d,d)
	e := a >> 2
	fmt.Printf("%d\t%b\n",e,e)

	
	fmt.Printf("%d\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t%b\n",mb,mb)
	fmt.Printf("%d\t%b\n",gb,gb)
}