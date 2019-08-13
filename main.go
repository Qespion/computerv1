package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mgutz/ansi"
)

func epur_zero(result []string) []string {
	result[0] = strings.Replace(result[0], "X^0", "1", -1)
	result[1] = strings.Replace(result[1], "X^0", "1", -1)
	return result
}

func check_degree(result string) bool {
	red := ansi.ColorFunc("red+")
	for i, rune := range result {
		if rune == 94 {
			if result[i+1] > 50 {
				fmt.Printf(red("▶ WARNING : The polynomial degree is stricly greater than 2, I can't solve! \n"))
				return false
			}
			if i+2 < len(result) {
				if result[i+2] >= 48 && result[i+2] <= 57 {
					fmt.Printf(red("▶ WARNING : The polynomial degree is stricly greater than 2, I can't solve.! \n"))
					return false
				}
			}
		}
	}
	return true
}

func reduced_form(args string) {
	result := strings.Split(args, "=")
	result = epur_zero(result)
	for len(result[1]) > 0 {
		result[1] = result[1][1:]
		fmt.Println(result[1])
	}
	return
}

func main() {
	if len(os.Args[1:]) < 1 {
		red := ansi.ColorFunc("red+")
		fmt.Printf(red("▶ WARNING : Please insert THINGS has argument! \n"))
		return
	}
	args := strings.Replace(os.Args[1:][0], " ", "", -1)
	if check_degree(args) == false {
		return
	}
	reduced_form(args)
	return
}
