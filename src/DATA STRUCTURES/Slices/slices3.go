package main
import "fmt"
//deleting fro slice
func main() {

	a := []int {2,3,4,5,6,7}
	a = append(a[:2],a[3:]...) //deleting 4
	fmt.Println(a)

}