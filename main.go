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

func main() {
	if len(os.Args[1:]) <= 1 {
		red := ansi.ColorFunc("red+")
		fmt.Printf(red("▶ WARNING : Please insert THINGS has argument! \n"))
		return
	}
	args := strings.Replace(os.Args[1:][0], " ", "", -1)
	result := strings.Split(args, "=")
	fmt.Println(result)
	result = epur_zero(result)
	fmt.Println(result)
	return
}
