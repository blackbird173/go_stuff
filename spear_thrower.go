package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to spear thrower the larger the input the father u throw")
	you := "|"
	terrian := "...................."
	var inp string
	fmt.Scan(&inp)
	x := 0
	if len(inp) <= len(terrian) {
		for x < len(inp) {
			fmt.Println(you + terrian)
			terrian = terrian[:x] + "-" + terrian[x+1:]
			if x > 0 {
				terrian = terrian[:x-1] + "." + terrian[x:]
			}
			x++
		}
		fmt.Println(you + terrian)
		terrian = terrian[:x-1] + "|" + terrian[x:]
		fmt.Println(you + terrian)
		fmt.Print("u got ", x, " points")
	} else {
		fmt.Println("the have passed the max throw length")
	}
}
