package main

import "fmt"

func main() {
    s := make([]string, 0, 6)
    m := make(map[string]string)
    str := new(string)
    m["m1"] = "m"
    m["m2"] = "mm"
    delete(m, "m2")
    s = append(s, "yeah", "sorry", "kid")
    rand_slice := make([]int, 3, 6)
    copy(rand_slice, []int{1, 2, 3})
    fmt.Println(m, s, str, rand_slice, len(s), cap(s))
    fmt.Println("make takes 3 args (slice , size, max size it can hold), new takes one arugment (type (can be int, string, bool, slice, array, etc)), delete takes two arugemnts a map and then part of the map u want to delete, append takes a slice as the first arg and the next args r the things u want to add on, copy takes a slice as the first arg and another slice as the next the second arugment copys the contents onto the first, cap returns the max size the arugment which is a slice can hold, len returns the size of something for example a slice")
}
