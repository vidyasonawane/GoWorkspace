package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {

	s := `[{"First":"vidya","Age":22},{"First":"kunal","Age":18},{"First":"vira","Age":28}]`
	fmt.Println("Json : ", s)
	var users []user

	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Data : ", users)

}
