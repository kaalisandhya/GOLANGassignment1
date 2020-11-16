1.package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 12; i++ {
		sum += i
	}
	fmt.Println(sum)
}
66

Program exited.
2.package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
1024

Program exited.
3.package main

import "fmt"

func main() {
	sum := 1
	for sum < 1200 {
		sum += sum
	}
	fmt.Println(sum)
}
2048

Program exited.
4.package main

func main() {
	for {
	}
}
Waiting for remote server...
Program exited: killed
5.package main
 
import (
    "fmt"
)
 
func main() {
    for i:=0; i<5; i++ {               
        fmt.Println("Hello golang")       
    }                               
}
$go run main.go
Hello golang
Hello golang
Hello golang
Hello golang
Hello golang