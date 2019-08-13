package main

import (
	"fmt"
	"os"
	"strings"
)

func epur_zero(result []string) []string {
	result[0] = strings.Replace(result[0], "X^0", "1", -1)
	result[1] = strings.Replace(result[1], "X^0", "1", -1)
	return result
}

func main() {
	args := strings.Replace(os.Args[1:][0], " ", "", -1)
	// arg1 := args[:strings.IndexByte(args, '=')]
	result := strings.Split(args, "=")
	fmt.Println(result)
	result = epur_zero(result)
	fmt.Println(result)
	return
}
