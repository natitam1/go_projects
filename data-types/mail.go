package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

// Function that uses the struct
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

func main() {
    // Create a struct value
    person1 := Person{Name: "Natnael", Age: 21}

    // Access struct fields
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)

    // Call a method on the struct
    person1.Greet()
}
