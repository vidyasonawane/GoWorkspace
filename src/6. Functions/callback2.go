package main

import "fmt"

func filter(numbers []int, cbfunc func(int) bool) []int {
	var xs []int
	for _, v := range numbers {
		if cbfunc(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
			return n > 1
		})

	fmt.Println(xs)
}
