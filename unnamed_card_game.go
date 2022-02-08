package main
import(
    "fmt"
    "math/rand"
    "time"
    "os"
)
func main() {
    rand.Seed(time.Now().UnixNano())
    u := []int{}
    e := []int{}
    deck := []int{}
    x := 52
    for z := 0; z <= 52; z++ {
        deck = append(deck, rand.Intn(9))
    }
    for i := 0; i <= 4; i++ {
        u = append(u, rand.Intn(9))
        e = append(e, rand.Intn(9))
    }
    zz := 0
    xx := 0
    z1 := 0
    x1 := 0
    for true {
        ranint := rand.Intn(9)
        fmt.Print("urs: ")
        fmt.Println(u)
        fmt.Print("his: ")
        fmt.Println(e)
        for i, v := range u {
            if i == 0 {

            }
            if v == zz {
                fmt.Println(z1)
                if zz == 0 {
                    z1++
                }
                if z1 >= 3 {
                    os.Exit(0)
                }
            }
        } 
        for x, d := range e {
            if x == 0 {

            }
            if d == xx {
                fmt.Println(x1)
                if xx == 0 {
                    x1++
                }
                if x1 >= 3 {
                    os.Exit(0)
                }
            }
        } 
        var inp string
        fmt.Scan(&inp)
        switch inp {
        case "draw":
            u = append(u, rand.Intn(9))
            if x == 0 {
                deck = []int{}
                for z := 0; z <= 52; z++ {
                    deck = append(deck, rand.Intn(9))
                    fmt.Print(z)
                    fmt.Println("cards left in the deck will reshuffle when it gets to 0")
                }
            }
        case "swap":
            fmt.Print("pick one of ur cards it starts at 0 and u have  ")
            fmt.Print(len(e))
            fmt.Println(" cards")
            var ucard int
            fmt.Scan(&ucard)
            fmt.Print("pick one of his cards it starts at 0 and he has ")
            fmt.Print(len(e))
            fmt.Println(" cards")
            var ecard int
            fmt.Scan(&ecard)
            z := u[ucard]
            u[ucard] = e[ecard]
            e[ecard] = z
        }
        if ranint >= 5 {
            e = append(e, rand.Intn(9))
        } else if ranint >= 7 {
            u[rand.Intn(9)] = e[rand.Intn(9)]
            e[rand.Intn(9)] = u[rand.Intn(9)]
        }
    }
}
