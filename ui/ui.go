package ui

import (
	"fmt"
)

func BoolPrompt(q string) bool {
	i := 0
	for i == 0 {
		var response int
		fmt.Println(q, " [y/n]")
		fmt.Printf("> ")
		fmt.Scanf("%c\n", &response)
		switch response {
		default:
			fmt.Println("Please answer 'y' or 'n'.")
		case 'y':
			return true
			i = 1
		case 'n':
			return false
			i = 1
		}
	}
	return false
}
func IntPrompt(q string) int {
	var response int
	i := 0
	for i == 0 {
		fmt.Println(q)
		fmt.Printf("> ")
		fmt.Scanf("%d\n", &response)
		if response >= 1 {
			i = 1
		}
	}
	return response
}
