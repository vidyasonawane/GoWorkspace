package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

//OR
// We can also use tags to compare json fields and struct fields
// type person struct {
// 	First string `json:"First"`
// 	Last  string `json:"Last"`
// 	Age   int    `json:"Age"`
// }

func main() {
	s := `[{"First":"vidya","Last":"sonawane","Age":22},{"First":"kunal","Last":"patil","Age":19}]`
	bs := []byte(s)
	fmt.Println(s)
	// fmt.Printf("%T\n",s)
	// fmt.Println(bs)
	// fmt.Printf("%T\n",bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("\nPerson number: ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
