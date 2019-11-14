package main
import "fmt"
//appending a slice to a slice
func main() {

	a := []int {2,3,4,5}
	b := []int{6,7,8,9}
	a = append(a,b...)
	fmt.Println(a)

}