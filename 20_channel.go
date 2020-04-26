package main

import (
    "fmt"
    "time"
)

func hello(done chan bool){
    fmt.Println("hello go routine is going to sleep")
    fmt.Println("hello world goroutine")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- true
}

func calculateSquares(num int, squareNum chan int){
    sum := 0
    for num != 0 {
        digit := num % 10
        sum += digit * digit
        num = num / 10
    }
    squareNum <- sum
}

func calculateCubes(num int, cubeNum chan int){
    sum := 0
    for num != 0 {
        digit := num % 10
        sum += digit * digit * digit
        num = num / 10
    }
    cubeNum <- sum
}

func main() {
    num := 123
    squareNum := make(chan int)
    cubeNum := make(chan int)
    go calculateSquares(num, squareNum)
    go calculateCubes(num, cubeNum)

    squares := <-squareNum
    cubes := <-cubeNum
    fmt.Println("this is result:")
    fmt.Println(squares + cubes)

    /*
    var a chan int
    if a == nil {
        fmt.Println("Channel a is nil, going to define it")
        a = make(chan int)
        fmt.Printf("Type of a is %T", a)
    }
    done := make(chan bool)
    go hello(done)
    <-done
    fmt.Println("main function")
    */
}
