package main

import (
    "fmt"
)


func changeLocal(num[5]int){
    num[0] = 0
    fmt.Println("inside", num)
}


func main(){
    var a [3] int
    a[0] = 23
    a[1] = 45
    a[2] = 56
    fmt.Println(a)

    var b = [3] int {1,2,4}
    fmt.Println(b)

    c := [3]int{12}
    fmt.Println(c)

    d := [...] int {1,3,5,7,8,2,4}
    fmt.Println(d)

    //数组是值类型 数组赋值给一个新的变量时，该变量会得到原始数组的一个副本 对新变量修改不会影响原数组

    e := [...] string {"USE", "China", "India", "Germany", "France"}
    f := e
    f[0] = "Singapore"
    fmt.Println(e)
    fmt.Println(f)


    num := [...] int {5,6,7,8,9}
    fmt.Println("before", num)
    changeLocal(num)
    fmt.Println("after", num)

    g := [...]float64{67.7, 89.5, 32, 78}
    fmt.Println("g length is ", len(g))


    //for circle travel
    for i:=0; i < len(g); i++ {
        fmt.Println(i, g[i])
    }

    for i, v := range g {
        fmt.Println(i, v)
    }




    //slice
    o := [5] int {76,77,78,79,80}
    //切片必须显示创建
    var p []int = o[1:4]
    fmt.Println(p)


    //切片自己不拥有任何数据 对切片的任何修改都会反映在底层数组上
    darr := [...] int {57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before", darr)
    for i := range dslice {
        dslice[i] ++
    }
    fmt.Println("array after ", darr)


    fruitarray := [...] string {"apple", "orange", "grape", "mango", "watermelon", "pineapple", "chikoo"}
    fruitslice := fruitarray[2:5]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))

    q := make([]int, 5, 5)
    fmt.Println(q)

    //追加切片元素
    cars := [] string {"car1", "car2", "car3"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
    cars = append(cars, "car4")
    fmt.Println("cars", cars, "has new length", len(cars), "and capacity", cap(cars))
}






























