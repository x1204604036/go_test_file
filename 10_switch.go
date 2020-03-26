package main

import (
    "fmt"
)


func getNumber()(int) {
    return 4*12
}


func main() {
    //finger := 4
    switch finger := 8; finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default:
        fmt.Println("incorrect finger")
    }
    //fmt.Println(finger)


    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u":
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }

    //number := getNumber()
    switch number := getNumber(); number > 0{
    case number < 50:
        fmt.Println(number)
        fallthrough //jump to next case 
    case number > 50:
        fmt.Println(number, "came here")
    }
}










