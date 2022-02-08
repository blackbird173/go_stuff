package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("welcome to the crab casino i assume ur here for the troll cards game")
	fmt.Println("here r the rules u can draw, swap, and show cards, but the dealer can only draw")
	file, _ := os.OpenFile("cardgamelogs.txt", os.O_CREATE|os.O_APPEND, 1337)
	rand.Seed(time.Now().UnixNano())
	u := []int{}
	d := []int{}
	for i := 0; i <= 10; i++ {
		u = append(u, rand.Intn(10))
		d = append(d, rand.Intn(10))
	}
	for true {
		m := rand.Intn(9)
		fmt.Print("urs: ")
		fmt.Println(u)
		fmt.Print("dealers: ")
		fmt.Println(d)
		var inp string
		fmt.Scan(&inp)
		switch inp {
		case "draw":
			u = append(u, rand.Intn(10))
		case "swap":
			fmt.Println("select a card pos starts at 0")
			var z int
			fmt.Scan(&z)
			x := u[z]
			fmt.Println("select a card pos starts at 0")
			var y int
			fmt.Scan(&y)
			u[z] = d[y]
			d[y] = x
			d[rand.Intn(10)] = x
		case "show":
			r := 0
			z := 0
			for i, v := range u {
				if i == 0 {

				}
				r += v
			}
			for x, f := range d {
				if x == 0 {

				}
				z += f
			}
			if r > z {
				file.Write([]byte("u win\n"))
				os.Exit(0)
			} else if z > r {
				fmt.Println([]byte("u lose\n"))
				os.Exit(0)
			}
		}
		if m >= 5 {
			d = append(d, rand.Intn(10))
		}
	}
}
