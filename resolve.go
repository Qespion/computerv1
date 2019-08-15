package main

import (
	"fmt"
	"math"
)

//getting x0 then x1 then x2
// delta = b^2 - 4ac
func resolve(tab [4]int) {
	var tab_float [3]float64
	tab_float[0] = float64(tab[0])
	tab_float[1] = float64(tab[1])
	tab_float[2] = float64(tab[2])
	fmt.Println(tab)
	delta_b2 := math.Pow(float64(tab[1]), 2)
	deltar_4ac := 4 * tab[2] * tab[0]
	delta := math.Pow(tab_float[1], 2) - 4*tab_float[2]*tab_float[0]
	fmt.Println(delta_b2, deltar_4ac)
	fmt.Println(delta)
	if delta == 0 || tab[2] == 0 {
		fmt.Printf("The solution is:\n%f\n", -tab_float[1]/2*tab_float[2])
	} else if delta > 0 {
		fmt.Printf("The solution is:\n%f\n", -tab_float[1]-math.Sqrt(delta)/2*tab_float[2])
		fmt.Printf("The solution is:\n%f\n", -tab_float[1]+math.Sqrt(delta)/2*tab_float[2])
	} else if delta < 0 {
		fmt.Printf("The solution is:\n%f + i * %f\n", (-tab_float[1]), math.Sqrt(-delta)/2*tab_float[2])
		fmt.Printf("The solution is:\n%f - i * %f\n", (-tab_float[1]), math.Sqrt(-delta)/2*tab_float[2])
	}

	return
}
