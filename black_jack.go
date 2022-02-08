package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("welcome to the crab casino i assume ur here for the blackjack game")
    eef, _ := os.OpenFile("crab_casino.txt", os.O_CREATE|os.O_APPEND, 1337)
    i := 0
    sl := []int{}
    dc := []int{}
    for i <= 3 {
        sl = append(sl, rand.Intn(7))
        dc = append(dc, rand.Intn(7))
        i++
    }
    var inp string
    for true {
        fmt.Scan(&inp)
        switch inp {
        case "draw":
            ra := rand.Intn(9)
            sl = append(sl, ra)
        case "show":
            d := 0
            y := 0
            for i, v := range sl {
                if i > 20 {
                }
                y += v
            }
            for x, z := range dc {
                if x > 20 {
                }
                d += z
            }
            fmt.Print("dealer's cards sum: ")
            fmt.Println(d)
            fmt.Print("ur cards sum: ")
            fmt.Println(y)
            if y > 21 && d > 21 {
                eef.Write([]byte("noone wins\n"))
                os.Exit(0)
            } else if 21-y < 21-d || y-21 > d-21 {
                eef.Write([]byte("u win\n"))
                os.Exit(0)
            } else if 21-y > 21-d || y-21 < d-21 {
                eef.Write([]byte("u lose\n"))
                os.Exit(0)
            } else if y == d {
                eef.Write([]byte("noone wins\n"))
                os.Exit(0)
            }
            if rand.Intn(10) > 5 {
                dc = append(dc, rand.Intn(7))
            }
        }
    }
}
