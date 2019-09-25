package main

import "fmt"

func main() {
	//using composite literals
	//almost same as array, only the size is omitted.
	fmt.Println("slice creation using composite literals")
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	//slicing a slice
	fmt.Println("slicing a slice")
	y := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(y)
	fmt.Println(y[:])
	fmt.Println(y[2:])
	fmt.Println(y[:4])  //all elements before index 4, excluding it
	fmt.Println(y[2:5]) // all elements between index 2(including) and 5(excluding)

	//append to a slice
	fmt.Println("append to a slice")
	z := []int{1, 2, 3}
	z = append(z, 4, 5, 6) //append should return
	fmt.Println(z)

	//append slice to slice
	fmt.Println("append slice to slice")
	a := []int{2, 4}
	b := []int{6, 8}
	a = append(a, b...)
	fmt.Println(a)

	//delete from slice
	fmt.Println("delete from slice")
	p := []int{1, 2, 3, 4, 5, 6, 7, 8}
	p = append(p[:3], p[3:]...) //delete 4
	fmt.Println(p)

	//slice using make built in function
	fmt.Println("slice using make built in function")
	s := make([]int, 2, 4)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	//s[2] =10 //index out of bound
	s = append(s, 10)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, 20)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, 30)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s)) //underlying array capacity becomes 8

	//multidimensional slices
	fmt.Println("multidimensional slices")
	s1 := []string{"vids", "kunal"}
	s2 := []string{"kamal", "vishnu"}
	ms := [][]string{s1, s2}
	fmt.Println(ms)

	ms2 := [][]string{{"sharayu", "teju"}, {"akanksha", "monika"}}
	fmt.Println(ms2)

}
