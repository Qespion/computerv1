package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"
)

func check_degree() bool {
	result := os.Args[1:][0]
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

func reducing_tab() [4]int {
	base_tab := strings.Fields(os.Args[1:][0])
	var tab [4]int
	var equal bool
	for i, nb := range base_tab {
		if nb == "=" {
			equal = true
		}
		stock, _ := strconv.Atoi(nb)
		if i != 0 && base_tab[i-1] == "-" {
			stock *= -1
		}
		if stock != 0 && i+1 < len(base_tab) && base_tab[i+1] == "*" {
			switch base_tab[i+2] {
			case "X^0":
				if equal {
					tab[0] -= stock
				} else {
					tab[0] += stock
				}
			case "X^1":
				if equal {
					tab[1] -= stock
				} else {
					tab[1] += stock
				}
			case "X^2":
				if equal {
					tab[2] -= stock
				} else {
					tab[2] += stock
				}
			}
		} else if stock != 0 {
			if equal {
				tab[3] -= stock
			} else {
				tab[3] += stock
			}
		}
		fmt.Println(stock)
		stock = 0
	}
	return tab
}

func render_reduced_form(tab [4]int) {
	degree := 0
	fmt.Printf("Reduced Form: ")
	if tab[0] != 0 {
		if tab[0] < 0 {
			fmt.Printf("- %d * X^0 ", tab[0]*-1)
		} else {
			fmt.Printf("%d * X^0 ", tab[0])
		}
	}
	if tab[1] != 0 {
		if tab[1] < 0 {
			fmt.Printf("- %d * X^1 ", tab[1]*-1)
		} else {
			fmt.Printf("+ %d * X^1 ", tab[1])
		}
		degree = 1
	}
	if tab[2] != 0 {
		if tab[2] < 0 {
			fmt.Printf("- %d * X^2 ", tab[2]*-1)
		} else {
			fmt.Printf("+ %d * X^2 ", tab[2])
		}
		degree = 2
	}
	fmt.Printf("= 0\n")
	fmt.Printf("Polynomial degree: %d\n", degree)
}

func main() {
	if len(os.Args[1:]) < 1 {
		red := ansi.ColorFunc("red+")
		fmt.Printf(red("▶ WARNING : Please insert THINGS as argument! \n"))
		return
	}
	if !check_degree() {
		return
	}
	tab := reducing_tab()
	render_reduced_form(tab)
	resolve(tab)
	/*
		Basic testcase:
		1 * X^0 + 2 * X^1 = - 1 * X^0 + 4 * X^1
	*/
	fmt.Println("ExpectedForm: 2 * X^0 - 2 * X^1 = 0 ")
	return
}
