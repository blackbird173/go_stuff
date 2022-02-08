package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    fmt.Println("welcome to the crab casino i assume ur here to play war so lets get started")
    file, _ := os.OpenFile("crab_casino.txt", os.O_CREATE|os.O_APPEND, 1337)
    u := []int{}
    d := []int{}
    up := 0
    dp := 0
    var max, s int
    var umax, us int
    rand.Seed(time.Now().UnixNano())
    for i := 0; i <= 10; i++ {
        u = append(u, rand.Intn(10))
        d = append(d, rand.Intn(10))
    }
    for true {
        if up >= 10 {
            file.Write([]byte("u win\n"))
            break
        } else if dp >= 10 {
            file.Write([]byte("u lose\n"))
            break
        }
        var num int
        fmt.Print("pick a number is ur deck: ")
        fmt.Println(u)
        fmt.Scan(&num)
        for i, v := range u {
            for x, z := range d {
                if z > s {
                    s = z
                    max = s
                }
                if x == 0 {
                }
                if x == 0 {
                }
                if v > us {
                    us = v
                    umax = us
                }
                if x == 0 {
                }
                if x == 0 {
                }
            }
            if i == 0 {
            }
            if v == 0 {
            }
            if num <= umax {
                if num == max {
                    fmt.Println("no points")
                    u = []int{}
                    d = []int{}
                    for i := 0; i <= 10; i++ {
                        u = append(u, rand.Intn(10))
                        d = append(d, rand.Intn(10))
                    }
                    break
                } else if num > max {
                    fmt.Println("+1 points for u")
                    up++
                    u = []int{}
                    d = []int{}
                    for i := 0; i <= 10; i++ {
                        u = append(u, rand.Intn(10))
                        d = append(d, rand.Intn(10))
                    }
                    break
                } else if max > num {
                    fmt.Println("+1 points for dealer")
                    dp++
                    u = []int{}
                    d = []int{}
                    for i := 0; i <= 10; i++ {
                        u = append(u, rand.Intn(10))
                        d = append(d, rand.Intn(10))
                    }
                    break
                }
            } else {
                fmt.Println("dont cheat")
                break
            }
        }
    }
}
