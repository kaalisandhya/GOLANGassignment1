1.package main 
  
import "fmt"
  
func main() { 
  
    var totalsum int = 8946 
    var number int = 19 
    var avg float32 
  
    
    avg = float32(totalsum) / float32(number) 
  
    
    fmt.Printf("Average = %f\n", avg) 
} 
Average = 470.842102

Program exited.
2.package main 
  
import ( 
    "fmt"
    "math/rand"
) 
  
// Main function 
func main() { 
  
    // Finding random numbers of float64 type 
    // Using Float64() function 
    res_1 := rand.Float64() 
    res_2 := rand.Float64() 
    res_3 := rand.Float64() 
  
    // Displaying the result 
    fmt.Println("Random Number 1: ", res_1) 
    fmt.Println("Random Number 2: ", res_2) 
    fmt.Println("Random Number 3: ", res_3) 
} 
Random Number 1:  0.6046602879796196
Random Number 2:  0.9405090880450124
Random Number 3:  0.6645600532184904
3.package main 
  
import ( 
    "fmt"
    "math/rand"
) 
  
 
func floatrandom(val1, val2 float64) float64 { 
    return val1 + val2 + rand.Float64() 
} 
  

func main() { 
  
    
    result1 := floatrandom(35, 56) 
    result2 := floatrandom(36, 50) 
    result3 := floatrandom(199, 700) 
  
    
    fmt.Println("Result 1: ", result1) 
    fmt.Println("Result 2: ", result2) 
    fmt.Println("Result 3: ", result3) 
} 
Result 1:  91.60466028797961
Result 2:  86.94050908804502
Result 3:  899.6645600532185
4.package main 
   
import ( 
    "fmt"
    "reflect"
) 
  
func main() { 
   
    var val chan int
   
   
    value := reflect.MakeChan(reflect.Indirect(reflect.ValueOf(&val)).Type(), 0) 
   
    fmt.Println("Value :", value) 
} 
Value : 0xc00005e060

Program exited.
5.package main 
   
import ( 
    "fmt"
    "reflect"
) 
   
func main() { 
   
    var str map[int]string  
       
    var strValue reflect.Value = reflect.ValueOf(&str) 
   
    indirectStr := reflect.Indirect(strValue) 
       
    //Use of Type() method 
   
    valueMap := reflect.MakeMap(indirectStr.Type()) 
       
    fmt.Printf("ValueMap is [%v] .", valueMap) 
   
} 
ValueMap is [map[]] .
Program exited.
