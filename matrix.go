package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    space := ""
    spaces := []string{}
    i := 0
    for i <= 119 {
        spaces = append(spaces, space)
        space = space + " "
        i++
    }
    x := 0
    clrs := []string{"\x1b[30m", "\x1b[31m", "\x1b[32m", "\x1b[33m", "\x1b[34m", "\x1b[35m", "\x1b[36m", "\x1b[37m", "\x1b[90m", "\x1b[91m", "\x1b[92m", "\x1b[93m", "\x1b[94m", "\x1b[95m", "\x1b[96m", "\x1b[97m"}
    chars := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
    rand.Seed(time.Now().UnixNano())
    for true {
        random1 := rand.Intn(119)
        if x >= 20 {
            x = 0
            fmt.Println("\033[0;0H")
        }
        for x <= 20 {
            random2 := rand.Intn(26)
            random3 := rand.Intn(16)
            fmt.Println(clrs[random3] + spaces[random1] + chars[random2])
            x++
            time.Sleep(1e+7)
        }
    }
}
