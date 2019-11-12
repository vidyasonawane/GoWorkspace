package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string //Fields should be in capital.
	Age   int
}

func main() {
	u1 := user{
		First: "vidya",
		Age:   22,
	}
	u2 := user{
		First: "kunal",
		Age:   18,
	}
	u3 := user{
		First: "vira",
		Age:   28,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
