1.package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(8))
}
$go run main.go
40320
2.package main

import "fmt"

func factorial(i int)int {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}
func main() { 
   var i int = 9
   fmt.Printf("Factorial of %d is %d", i, factorial(i))
}
$go run main.go
Factorial of 9 is 362880
3.package main

import "fmt"

func fibonaci(i int) (ret int) {
   if i == 0 {
      return 0
   }
   if i == 1 {
      return 1
   }
   return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
   var i int
   for i = 0; i < 20; i++ {
      fmt.Printf("%d ", fibonaci(i))
   }
}
$go run main.go
0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 
4.package main 
  
import ( 
    "fmt"
) 
func factorial_calc(number int) int { 
  
    if number == 0 || number == 1 { 
        return 1 
    } 
      
    if number < 0 { 
        fmt.Println("Invalid number") 
        return -1 
    } 
    return number*factorial_calc(number - 1) 
} 
  
func main() { 
  
    answer1 := factorial_calc(0) 
    fmt.Println(answer1, "\n") 
         
    answer2 := factorial_calc(5) 
    fmt.Println(answer2, "\n") 
        
    answer3 := factorial_calc(-1) 
    fmt.Println(answer3, "\n") 
      
    answer4 := factorial_calc(10) 
    fmt.Println(answer4, "\n") 
} 
$go run main.go
1 

120 

Invalid number
-1 

3628800 
5.package main  
  import (  
 "fmt"  
)  
  func recursivefunction() {  
 fmt.Println("welcome")  
 recursivefunction()  
}  
  func main() {  
 recursivefunction()  
}  
recursion
recursion
recursion
recursion
recursion
recursion
........
infinte times
6.package main  
  
import (  
 "fmt"  
)  
  
var count = 0  
  
func recursivemethod() {  
 count++  
 if count <= 15 {  
  fmt.Println(count)  
  recursivemethod()  
 }  
}  
  
func main() {  
 recursivemethod()  
}  
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