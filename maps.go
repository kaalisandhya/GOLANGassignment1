1.package main
 
import "fmt"
 
var employee = map[string]int{"Mark": 10, "Sandy": 20}
 
func main() {
	fmt.Println(employee)
}
$go run main.go
map[Mark:10 Sandy:20]
2.package main
 
import "fmt"
 
func main() {
	var employee = map[string]int{}
	fmt.Println(employee)        // map[]
	fmt.Printf("%T\n", employee) // map[string]int
}
$go run main.go
map[]
map[string]int
3.package main
 
import "fmt"
 
func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 89
	employee["Sandy"] = 90
	fmt.Println(employee)
 
	employeeList := make(map[string]int)
	employeeList["Mark"] = 78
	employeeList["Sandy"] = 78
	fmt.Println(employeeList)
}
$go run main.go
map[Mark:89 Sandy:90]
map[Mark:78 Sandy:78]
4.package main
 
import "fmt"
 
func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 78
	employee["Sandy"] = 90
 
	// Empty Map
	employeeList := make(map[string]int)
 
	fmt.Println(len(employee))     // 2
	fmt.Println(len(employeeList)) // 0
}
$go run main.go
2
0
5.package main
 
import "fmt"
 
func main() {
	var employee = map[string]int{"Mark": 78, "Sandy": 28}
 
	fmt.Println(employee["Mark"])
}
$go run main.go
78
5.package main
 
import "fmt"
 
func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map
 
	employee["Rocky"] = 30 // Add element
	employee["Josef"] = 40
 
	fmt.Println(employee)
}
$go run main.go
map[Mark:10 Sandy:20]
map[poojitha:40 Mark:10 Sandy:20 prajwala:30]
6.package main
 
import "fmt"
 
func main() {
	var student = map[string]int{"Mark": 10, "Sandy": 90}
	fmt.Println(student) // Initial Map
 
	student["Mark"] = 90 // Edit item
	fmt.Println(student)
}
$go run main.go
map[Mark:10 Sandy:90]
map[Mark:90 Sandy:90]
7.package main
 
import "fmt"
 
func main() {
	var student = make(map[string]int)
	student["prajwala"] = 10
	student["Sandya"] = 20
	student["sujatha"] = 30
	student["panathi"] = 40
 
	fmt.Println(student)
 
	delete(student, "prajwala")
	fmt.Println(student)
}
$go run main.go
map[prajwala:10 Sandya:20 sujatha:30 panathi:40]
map[Sandya:20 sujatha:30 panathi:40]
8.package main
 
import "fmt"
 
func main() {
    var student = map[string]int{"sandhya": 11, "sujini": 22,
        "sujatha": 33, "raji": 44, "sushanth": 55}
    for key, element := range student {
        fmt.Println("Key:", key, "=>", "Element:", element)
    }
}
$go run main.go
Key: sujini => Element: 22
Key: sujatha => Element: 33
Key: raji => Element: 44
Key: sushanth => Element: 55
Key: sandhya => Element: 11