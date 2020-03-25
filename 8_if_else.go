package main

import (
    "fmt"
)

func main() {
    num := 10
    if num % 2 == 0 {
        fmt.Println("the number is even")
    } else {
        fmt.Println("the number is odd")
    }


    //the number only works below 
    if num1 := 7;  num1 % 2 == 0 {
        fmt.Println("the number is even")
    } else {
        fmt.Println("the number is odd")
    }

    //fmt.Println(num1) error
}

