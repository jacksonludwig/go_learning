package main

import (
    "fmt"
    "math"
)

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

    sum_any(1, 2, 3, 4)
}

func plus(a int, b int) int {
    return a + b
}

func sum_any(nums ...int) {
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func zeroval(ival int) {
    ival = 0
}
func zeroptr(iptr *int) {
    *iptr = 0
}


type person struct {
    name string
    age int
}
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

