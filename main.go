package main

import (
    "fmt"
    "package"
    "package2"
    "package3"
    "package4"
    "package5"
   
)

func main() {
   
    emp1 := &package.Employee1{}
    emp2 := &package2.Employee2{}
    emp3 := &package3.Employee3{}
    emp4 := &package4.Employee4{}
    emp5 := &package5.Employee5{}	
    
    emp1.SetPosition("Manager")
    emp1.SetSalary(50000)
    emp1.SetAddress("123 Main St")

    emp2.SetPosition("Engineer")
    emp2.SetSalary(60000)
    emp2.SetAddress("456 Elm St")
    
    emp3.SetPosition("Manager")
    emp3.SetSalary(50000)
    emp3.SetAddress("123 Main St")
    
    emp4.SetPosition("Manager")
    emp4.SetSalary(50000)
    emp4.SetAddress("123 Main St")

    emp5.SetPosition("Manager")
    emp5.SetSalary(50000)
    emp5.SetAddress("123 Main St")

    fmt.Printf("Employee 1 - Position: %s, Salary: %.2f, Address: %s\n", emp1.GetPosition(), emp1.GetSalary(), emp1.GetAddress())
    fmt.Printf("Employee 2 - Position: %s, Salary: %.2f, Address: %s\n", emp2.GetPosition(), emp2.GetSalary(), emp2.GetAddress())
    fmt.Printf("Employee 3 - Position: %s, Salary: %.2f, Address: %s\n", emp3.GetPosition(), emp3.GetSalary(), emp3.GetAddress())
    fmt.Printf("Employee 4 - Position: %s, Salary: %.2f, Address: %s\n", emp4.GetPosition(), emp4.GetSalary(), emp4.GetAddress())
    fmt.Printf("Employee 5 - Position: %s, Salary: %.2f, Address: %s\n", emp5.GetPosition(), emp5.GetSalary(), emp5.GetAddress())



}