package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	//adding new record
	m["sonawane_vidya"] = []string{"Piyu bole", "man mandira"}

	for i, m2 := range m {
		fmt.Println("Record of ", i)
		for j, val := range m2 {
			fmt.Println(j, val)
		}
	}

}
