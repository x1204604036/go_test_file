package main

import (
    "fmt"
    "math"
)

func main() {
    var age int //如果不赋值会自动初始化赋值为0 
    fmt.Println("ths age is ", age)

    age = 23
    fmt.Println("ths age is ", age)

    age = 53
    fmt.Println("the age is ", age)


    var age2 int = 29
    fmt.Println("ths age2 is ", age2)

    //if there is a init value, we can omit the type 
    var age3 = 56
    fmt.Println("the age3 is ", age3)


    //declare multi variables at the same time
    var width, height int = 100, 50
    var width1, height1 int
    var width2, height2 = 100, 50
    fmt.Println("this is width, height ")
    fmt.Println(width, height)
    fmt.Println(width1, height1)
    fmt.Println(width2, height2)


    //declare different type of variables
    var (
        name = "hunter"
        age4 = 23
        height4 int
    )
    fmt.Println("name is ", name, " the age is ", age4,"  height is ", height4 )


    //short hand decaration 
    name5, age5 := "hunter", 56
    fmt.Println("name5 is", name5, "age is", age5)


    //short hand decaration demand at least one variable is not decared on the left
    a, b := 1, 2
    b, c := 3,4
    b, c = 5,6
    fmt.Println("ths value of a, b, c is", a, b, c)


    a1, b1 := 145.0, 356.0
    c1 := math.Min(a1, b1)
    fmt.Println("a1, b1, c1 is ", a1, b1, c1)

    //Go is a langeuage called Strongly Typed, so it's not allowed for go to change type like 
    //a1 := 2
    //a1 = "string"
    //it's not allowed above

}










