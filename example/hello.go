package main

import "fmt"

const s string = "constant"

func main() {
    a := "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // continue, break, etc. all works

    var x [5]int
    x[4] = 100
    fmt.Println(x)

    slice := make([]string, 3)
    fmt.Println(slice)

    slice[0] = "a"
    slice[1] = "b"
    slice[2] = "c"
    fmt.Println(slice)

    m := make(map[string]int) // string key, int value
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println(m)

    delete(m, "k2")
    _, pair := m["k2"]
    fmt.Println(pair)
}
