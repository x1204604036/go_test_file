package main

import (
    "fmt"
    "math"
)

func main() {
    const a = 55
    //a = 89
    fmt.Println(a)

    //the value of a const will be confirmed
    //so we cannot assign the value to a const
    //when a function is called like...
    var a1 = math.Sqrt(4) //OK
    fmt.Println(a1)
    //const b1 = math.Sqrt(4) // not allowed
    //fmt.Println(b1)
}
