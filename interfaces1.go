1.package main

import (  
    "fmt"
)


type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string


func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

func main() {  
    name := MyString("old mahabalipuram")
    var v VowelsFinder
    v = name // possible since MyString implements VowelsFinder
    fmt.Printf("Vowels are %c", v.FindVowels())

}
Vowels are [o a a a i u a]
Program exited.
2.package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId    int
    basicpay int
}


func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}


func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{
        empId:    111,
        basicpay: 8000,
        pf:       29,
    }
    pemp2 := Permanent{
        empId:    112,
        basicpay: 9000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    113,
        basicpay: 11000,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}
Total Expense Per Month $28059
Program exited.
3.package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId    int
    basicpay int
}

type Freelancer struct {  
    empId       int
    ratePerHour int
    totalHours  int
}


func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}


func (c Contract) CalculateSalary() int {  
    return c.basicpay
}


func (f Freelancer) CalculateSalary() int {  
    return f.ratePerHour * f.totalHours
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{
        empId:    11,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    12,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    13,
        basicpay: 3000,
    }
    freelancer1 := Freelancer{
        empId:       14,
        ratePerHour: 70,
        totalHours:  120,
    }
    freelancer2 := Freelancer{
        empId:       15,
        ratePerHour: 100,
        totalHours:  1000,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
    totalExpense(employees)
}
Total Expense Per Month $122450
Program exited.
4.package main

import (  
    "fmt"
)

type Worker interface {  
    Work()
}

type Person struct {  
    name string
}

func (p Person) Work() {  
    fmt.Println(p.name, "is very dedicated")
}

func describe(w Worker) {  
    fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {  
    p := Person{
        name: "sandhya",
    }
    var w Worker = p
    describe(w)
    w.Work()
}
Interface type main.Person value {sandhya}
sandhya is very dedicated

Program exited.
5.package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello golang class"
    describe(s)
    i := 50
    describe(i)
    strt := struct {
        name string
    }{
        name: "sandhya kaali",
    }
    describe(strt)
}
Type = string, value = Hello golang class
Type = int, value = 50
Type = struct { name string }, value = {sandhya kaali}
