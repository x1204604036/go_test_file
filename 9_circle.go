package main

import (
    "fmt"
)


func main() {
    for i := 1; i <= 10; {
        if i >= 5 {
            break
        }
        //if i % 2 == 0 {
        //    continue
        //}
        fmt.Println(i)
        i += 1
    }
    //fmt.Println(i)


    //multi variables

    for no, i := 10; i <= 10 && no <= 19; i, no = ii + 1, no + 1 {
        fmr.Println(no, i, no*i)
    }

}











