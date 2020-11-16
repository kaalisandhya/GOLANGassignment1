1.package main 
  
import "fmt"
  
func myfunc(ch chan int) { 
  
    fmt.Println(239 + <-ch) 
} 
func main() { 
    fmt.Println("start Main method") 
   
    ch := make(chan int) 
  
    go myfunc(ch) 
    ch <- 23 
    fmt.Println("End Main method") 
} 
start Main method
262
End Main method
2.package main 
  
import "fmt"
  

func myfun(mychnl chan string) { 
  
    for v := 0; v < 4; v++ { 
        mychnl <- "sandhya kaali"
    } 
    close(mychnl) 
} 
  

func main() { 
  
    // Creating a channel 
    c := make(chan string) 
  
    // calling Goroutine 
    go myfun(c) 
  
    
    for { 
        res, ok := <-c 
        if ok == false { 
            fmt.Println("Channel Close ", ok) 
            break
        } 
        fmt.Println("Channel Open ", res, ok) 
    } 
} 
Channel Open  sandhya kaali true
Channel Open  sandhya kaali true
Channel Open  sandhya kaali true
Channel Open  sandhya kaali true