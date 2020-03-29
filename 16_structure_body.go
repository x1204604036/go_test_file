package main

import (
    "fmt"
    "structs/computer"
)


type Employee struct {
    firstName, lastName string
    age, salary         int
}


type Person struct {
    string
    int
}


type Address struct {
    city, state string
    age int
}

type Person1 struct {
    name string
    age int
    address Address
}

type Person2 struct{
    name string
    age int
    Address
}


type name struct {
    firstName string
    lastName string
}


func main(){
    emp1 := Employee{
        firstName: "Sam",
        age:        25,
        salary:     500,
        lastName:   "Anderson",
    }
    emp2 := Employee{"Thomas", "Paul", 29, 800}
    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)

    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName:      "Andreah",
        lastName:       "Nikola",
        age:            31,
        salary:         5000,
    }
    fmt.Println("Employee 3", emp3)

    var emp4 Employee
    fmt.Println("Employee 4", emp4)

    emp5 := Employee{
        firstName: "Johm",
        lastName:  "Paul",
    }
    fmt.Println("Employee 5", emp5)

    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d", emp6.salary)

    emp7 := Employee{
        firstName: "Hunter",
        age:       23,
        }
    emp7.lastName = "Xiong"
    emp7.salary = 16000
    fmt.Println(emp7)

    emp8 := &Employee{"Sam", "Anderson", 55, 8000}
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Last Name:", (*emp8).lastName)
    fmt.Println("Age:", emp8.age)


    p := Person{"Naveen", 50}
    fmt.Println(p)

    var p1 Person
    p1.string = "naveen"
    p1.int = 60
    fmt.Println(p1)

    var p2 Person1
    p2.name = "Naveen"
    p2.age = 50
    p2.address = Address{
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p2.name)
    fmt.Println("Age:", p2.age)
    fmt.Println("City:", p2.address.city)
    fmt.Println("State:", p2.address.state)

    var p3 Person2
    p3.name = "hunter"
    p3.age = 49
    p3.Address = Address{
        city: "Chicago",
        state: "Illinois",
        age: 50,
    }
    fmt.Println("........")
    fmt.Println("name:", p3.name)
    fmt.Println("age:", p3.age)
    fmt.Println("city:", p3.city)
    fmt.Println("state:", p3.state)

    var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 5000
    fmt.Println("Spec:", spec)

    name1 := name{"Steve", "Jobs"}
    name2 := name{"Steve", "Jobs"}
    if name1 == name2{
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name are not equal")
    }
}










