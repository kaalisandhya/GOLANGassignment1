package main    
import ( 
    "fmt"
    "encoding/json"
)   
type Human struct{ 
    
    Name string 
    Age int
    Address string
    
}   
func main() { 
    human1 := Human{"nikitha",23, "kodur"} 
       
    human_enc, err := json.Marshal(human1) 
       
    if err != nil { 
           
        
        fmt.Println(err) 
    } 
       
 $go run main.go
{"Name":"nikitha","Age":23,"Address":"kodur"}

[{"Name":"mahitha","Age":23,"Address":"mumbai"},{"Name":"karishma","Age":25,"Address":"Pune"},{"Name":"thulasi","Age":26,"Address":"Bangalore"}]   
    fmt.Println(string(human_enc)) 
       
    
    human2 := []Human{ 
        {Name: "mahitha", Age: 23, Address: "mumbai"},
        {Name: "karishma", Age: 25, Address: "Pune"}, 
        {Name: "thulasi", Age: 26, Address: "Bangalore"}, 
    } 
       
    
    human2_enc, err := json.Marshal(human2) 
        
        if err != nil { 
       
        
            fmt.Println(err) 
        } 
           
    
    fmt.Println() 
        fmt.Println(string(human2_enc)) 
} 
package main  
import "encoding/json"  
import "fmt"    
func main() {  
    bolType, _ := json.Marshal(false) //boolean Value  
    fmt.Println(string(bolType))  
    intType, _ := json.Marshal(78) // integer value  
    fmt.Println(string(intType))  
    fltType, _ := json.Marshal(89.7) //float value  
    fmt.Println(string(fltType))  
   
}  
$go run main.go
false
78
89.7