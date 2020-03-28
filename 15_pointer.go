package main

import (
    "fmt"
)


func change(val *int){
    *val = 55
}


func modify(sls []int){
    sls[0] = 90
}


func main() {
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is ", a)

    a1 := 25
    var b1 *int
    if b1 == nil {
        fmt.Println("b is ", b1)
        b1 = &a1
        fmt.Println("b after initialization is", b1)
    }

    b2 := 255
    a2 := &b2
    fmt.Println("address of b is ", a2)
    fmt.Println("value of b is ", *a2)

    b3 := 255
    a3 := &b3
    fmt.Println("address of b is ", a3)
    fmt.Println("address of b is ", *a3)
    *a3 ++
    b3 ++
    fmt.Println("b3 value is ", b3)


    //pointer to function
    a4 := 58
    fmt.Println("value of a before function call is ", a4)
    b4 := &a4
    change(b4)
    fmt.Println("value of a after function call is ", a4)

    //传递数组的切片 而不是数组的指针
    a5 := [3]int{89,90,91}
    modify(a5[:])
    fmt.Println(a5)


    //go 不支持指针运算
}











