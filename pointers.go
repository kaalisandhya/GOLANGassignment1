1.package main

import (  
    "fmt"
)

func main() {  
    b := 200
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is", a)
}
Type of a is *int
address of b is 0xc000100010

Program exited.
2.package main

import (  
    "fmt"
)
func main() {  
    a := 20
    var b *int
    if b == nil {
        fmt.Println("b is", b)
        b = &a
        fmt.Println("b after initialization is", b)
    }
}
b is <nil>
b after initialization is 0xc00002c008

Program exited.
3.package main

import (  
    "fmt"
)

func main() {  
    size := new(int)
    fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
    *size = 85
    fmt.Println("New size value is", *size)
}
Size value is 0, type is *int, address is 0xc00002c008
New size value is 80

Program exited.
4.package main

import (  
    "fmt"
)

func main() {  
    b := 205
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)
}
address of b is 0xc00002c008
value of b is 205
new value of b is 206

