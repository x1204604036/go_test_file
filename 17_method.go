package main

import (
    "fmt"
)

type Employee struct {
    name string
    salary int
    currency string
}

type Employee1 struct {
    name string
    age int
}

//displaySalary() 方法将Employee作为接收器类型
func (e Employee) displaySalary() {
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

//使用值接收器的方法
func (e Employee1) changeName(newName string){
    e.name = newName
}

//使用指针接收器的方法
func (e *Employee1) changeAge(newAge int){
    e.age = newAge
}

func main() {
    emp1 := Employee {
        name: "Sam Adolf",
        salary: 5000,
        currency: "$",
    }
    emp1.displaySalary()

    e := Employee1{
        name: "Mark Andrew",
        age: 50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Micheal Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    (&e).changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)

}








