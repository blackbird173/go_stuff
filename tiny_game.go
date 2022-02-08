package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    file, _ := os.OpenFile("g_logs.txt", os.O_CREATE|os.O_APPEND, 1337)
    x := 1
    rand.Seed(time.Now().UnixNano())
    dude := "|"
    line := ".........."
    goal := "{}"
    for true {
        num := rand.Intn(20)
        if len(line) == 0 {
            file.Write([]byte("u win \n"))
            os.Exit(0)
        }
        if line[0:1] == "<" && len(dude) == 2 {
            line = line[:0] + "" + line[1:]
            dude = dude[:1] + "" + dude[2:]
            x = 1
        } else if len(dude) == 1 && line[0:1] == "<" {
            file.Write([]byte("u lose \n"))
            os.Exit(0)
        }
        if num >= 10 {
            if x > 1 && x < len(line) {
                line = line[:len(line)-x-1] + "." + line[len(line)-x+1:]
            }
            if x < len(line) {
                line = line[:len(line)-x] + "<" + line[len(line)-x:]
            }
            x++
        }
        var inp string
        fmt.Println(dude + line + goal)
        fmt.Scan(&inp)
        switch inp {
        case "e":
            line = line + "."
        case "d":
            line = line[:0] + "" + line[1:]
        case "s":
            if string(dude[len(dude)-1:]) != "/" {
                dude = dude + "/"
            }
        default:
            fmt.Println("pls select (e, d, or s)")
        }
    }
}
