1.package main

import (  
    "fmt"
)

func change(val *int) {  
    *val = 55
}
func main() {  
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
value of a before function call is 58
value of a after function call is 55

Program exited.
2.package main

import (  
    "fmt"
)

func hello() *int {  
    i := 9
    return &i
}
func main() {  
    d := hello()
    fmt.Println("Value of d", *d)
}
Value of d 9

Program exited.
3.package main

import (  
    "fmt"
)

func modify(arr *[3]int) {  
    (*arr)[0] = 78
}

func main() {  
    a := [3]int{88,99,98}
    modify(&a)
    fmt.Println(a)
}
[78 99 98]

Program exited.
4.package main

import (  
    "fmt"
)

func modify(arr *[3]int) {  
    arr[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
[90 90 91]

Program exited.
5.package main

import (  
    "fmt"
)

func modify(sls []int) {  
    sls[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(a[:])
    fmt.Println(a)
}
[90 90 91]

Program exited.