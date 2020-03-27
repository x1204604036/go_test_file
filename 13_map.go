package main

import (
    "fmt"
)

func main() {
    personSalary := make(map[string]int)
    personSalary["steve"] = 12000
    personSalary["hunter"] = 15000
    personSalary["mike"] = 9000
    fmt.Println(personSalary)

    personSalary2 := map[string] int {
        "steve": 12000,
        "hunter": 15000,
        "jamie": 9000,
    }
    personSalary2["mike"] = 8000
    fmt.Println(personSalary2)

    employee := "hunter"
    fmt.Println("Salary of ", employee, "is", personSalary2[employee])
    fmt.Println("Salary of ", employee, "is", personSalary2["steve"])
    fmt.Println("Salary of ", "jack", "is", personSalary2["jack"])

    name := "sim"
    value, ok := personSalary[name]
    if ok == true {
        fmt.Println("Salary of ",name, "is", value)
    } else {
        fmt.Println(name, "not found")
    }
    for k, v := range personSalary {
        fmt.Println(k, v)
    }

    //fmt.Println("before deletion", personSalary)
    //delete(personSalary, "hunter")
    //fmt.Println("after deletion", personSalary)
    //delete(personSalary, "hunter")
    //fmt.Println("length of personSalary", len(personSalary))

    //map is 引用类型, map 赋值为一个新变量的时候，指向同一个内部数据结构，改变一个 影响到另一个
    fmt.Println("original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["hunter"] = 1
    fmt.Println("changed person salary", personSalary)

    //作为函数参数传递时 也会发生同样的情况 map 的任何修改 对于外部的调用都是可见的
    // == 只能用来检查 map 是否为 nil

}











