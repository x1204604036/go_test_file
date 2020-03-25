package main

import (
    "fmt"
    "math"
)

func main() {
    const a = 55
    //a = 89 //not allowed
    fmt.Println(a)

    b := math.Sqrt(4)
    fmt.Println(b)
    //const b1 = math.Sqrt(4) //not allowed

    //字符串常量是无类型的
    const hello = "Hello World"
    fmt.Println(hello)

    var name = "hunter"
    fmt.Println("type %T value %v", name, name)

    //bool const
    const trueConst = true
    fmt.Println(trueConst)


    const a2 = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println(intVar)
    fmt.Println(int32Var)
    fmt.Println(float64Var)
    fmt.Println(complex64Var)


    var i = 5
    var f = 5.6
    var c = 5 + 6i
    fmt.Println(i, f, c)


    var a4 = 5.9 / 8
    var b5 float64 = 5.6
    var a5 int = 7
    var c5 = b5 / float64(a5)
    fmt.Println(a4, a5, b5, c5)

}












