package main

import (
    "fmt"
)


func calculateBill(price, no int) (totalPrice, diff int) {
    //var totalPrice = price * no
    //var diff = price - no
    //return totalPrice, diff
    totalPrice = price * no
    diff = price - no
    return
}


func main() {
    //var price int = 3
    //var no int = 4
    price, no := 3, 4
    //totalPrice, diff := calculateBill(price, no)
    totalPrice, _ := calculateBill(price, no)
    fmt.Println(totalPrice)
}
