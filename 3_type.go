package main

import (
    "fmt"
    "unsafe"
)

//bool 
//number
// int8 int16 int32 int64 int
// uint8 uint16 uint32 uint64 int
// float32 float64
// complex64 complex128
// byte
// rune
//string

func main() {
    a := true
    b := false
    fmt.Println("a:", a, "b:", b)
    c := a&&b
    fmt.Println("c:", c)
    d := a||b
    fmt.Println("d:", d)

    var a1 int = 89
    b1 := 34
    fmt.Println("the a and b is ", a1, b1)

    fmt.Println("type of a is %T, size of a is %d", a1, unsafe.Sizeof(a1))
    fmt.Println("type of b is %T, size of b is %d", b1, unsafe.Sizeof(b1))


    //float introduction
    a2, b2 := 5.67, 8.97
    fmt.Println("type of a %T b %T\n", a2, b2)
    sum := a2 + b2
    diff := a2 - b2
    fmt.Println("sum", sum, "diff", diff)

    no1, no2 := 56, 89
    fmt.Println("sum", no1+no2, "diff", no1-no2)



    //string



    //type change
    i := 55
    j := 34.7
    //it's not allowed below
    //sum := i+j
    fmt.Println(i, j)
    fmt.Println(i+int(j))

    i1 := 56
    //j1 := float64(i1)
    var j1 float64 = float64(i1)
    fmt.Println(i1, j1)
}












