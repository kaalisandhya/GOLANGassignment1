1.package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    s := i.(int) 
    fmt.Println(s)
}
func main() {  
    var s interface{} = 50
    assert(s)
}
50

Program exited.
2.package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello sandhya"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "sandhya kaali",
    }
    describe(strt)
}
Type = string, value = Hello sandhya
Type = int, value = 55
Type = struct { name string }, value = {sandhya kaali}
3.package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    s := i.(int) 
    fmt.Println(s)
}
func main() {  
    var s interface{} = 59
    assert(s)
}
59

Program exited.
4.package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
func main() {  
    var s interface{} = 57
    assert(s)
    var i interface{} = "Sandhya kaali"
    assert(i)
}
50 true
0 false

Program exited.
5.package main

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("sandhya")
    findType(77)
    findType(89.98)
}
I am a string and my value is sandhya
I am an int and my value is 77
Unknown type
6.package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    findType("sandhya")
    p := Person{
        name: "sandhya kaali",
        age:  23,
    }
    findType(p)
}
unknown type
sandhya kaali is 23 years old
Program exited.