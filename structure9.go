package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}

type Person struct {  
    name    string
    age     int
    address Address
}

func main() {  
    p := Person{
        name: "sandhya",
        age:  50,
        address: Address{
            city:  "kadapa",
            state: "andhrapradesh",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("State:", p.address.state)
}
$go run main.go
Name: sandhya
Age: 50
City: kadapa
State: andhrapradesh