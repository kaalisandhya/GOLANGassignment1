package main
import "errors"
import "fmt"
import "math"
func Sqrt(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value), nil
}
func main() {
   result, err:= Sqrt(-1)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
   
   result, err = Sqrt(789)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
}
$go run main.go
Math: negative number passed to Sqrt
28.089143810376278
2.package main
import (
   "database/sql"
   "fmt"
)

func foo() error {
   return sql.ErrNoRows
}

func bar() error {
   return foo()
}

func main() {
   err := bar()
   if err != nil {
      fmt.Printf("got err, %+v\n", err)
   }
}
$go run main.go
got err, sql: no rows in result set