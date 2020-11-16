1.package main
 
import (
    "fmt"
)
 
func main() {
    i := 0
    for i<5 {
        fmt.Println(i)       
        i += 2              
    }
}
$go run main.go
0
2
4
2.package main
 
import (
    "fmt"
)
 
func main() {
 
    var items []int = []int{1, 2, 3, 4,5,6,7}
    for i, v := range items {
        fmt.Println(i, v)
    }
 }
$go run main.go
0 1
1 2
2 3
3 4
4 5
5 6
6 7
3.package main
 
import (
    "fmt"
)
 
func main() {
 
    var items []int = []int{1, 2, 3, 4, 5}
    for _, v := range items {
        fmt.Println(v)
    }
 
}
$go run main.go
1
2
3
4
5
4.package main
 
import (
    "fmt"
)
 
func main() {
 
    var m map[int]string = map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
        4:"four",
    }
     
    for k, v := range m {
        fmt.Println(k, v)
    }
 }
$go run main.go
1 One
2 Two
3 Three
4 four
5.package main
 
import (
    "fmt"
)
 
func main() {
    for i := 0; i<2; i++ {
        for j := 0; j<2; j++ {
            fmt.Println(i, j)
        } 
    }

}
$go run main.go
0 0
0 1
1 0
1 1
6.package main

import "fmt"

func main() {
     i := 1
     max := 20

    for i < max {
    fmt.Println(i)
  	 i += 1
     }
}
$go run main.go
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19	