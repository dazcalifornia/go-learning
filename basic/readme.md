# Go Programming Language Guide

This guide introduces fundamental concepts in Go programming through examples and explanations. It's designed as a reference for beginners learning the language.

## Table of Contents

1. [Setting Up a Go Program](#setting-up-a-go-program)
2. [Variables and Data Types](#variables-and-data-types)
3. [Constants](#constants)
4. [Arrays](#arrays)
5. [Slices](#slices)
6. [Maps](#maps)
7. [Structs](#structs)
8. [Functions](#functions)
9. [User Input](#user-input)
10. [Control Structures](#control-structures)
11. [Cleaned Example Code](#cleaned-example-code)

## Setting Up a Go Program

Every Go program starts with a package declaration and import statements:

```go
package main

import (
    "fmt"
)
```

The `main` package is special - it defines a standalone executable program, not a library. The `main` function in this package is the entry point of the program.

## Variables and Data Types

Go is statically typed. Variables can be declared in several ways:

### Explicit Declaration with Initialization

```go
var name string = "gopher"
var age int = 25
var height float32 = 5.9
var isActive bool = true
```

### Type Inference

Go can infer types from initialized values:

```go
name := "gopher"  // string
count := 10       // int
pi := 3.14        // float64
isValid := true   // bool
```

### Multiple Declaration

```go
var x, y int
var name, email string
```

### Basic Data Types

- **Numeric types**: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64
- **String type**: string
- **Boolean type**: bool

## Constants

Constants are values that cannot be changed during execution:

```go
const Pi = 3.14159
const MaxConnections = 100
const AppName string = "GoApp"
```

## Arrays

Arrays have a fixed size determined at declaration:

```go
// Declare with specific size and values
numbers := [6]int{10, 11, 12, 13, 14, 15}

// Let Go count the elements
days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri"}
```

## Slices

Slices are more flexible than arrays and are used more frequently:

```go
// Empty slice
emptySlice := []int{}

// Initialized slice
words := []string{"Go", "Slices", "Are", "Powerful"}

// Slice from an array
arr := [5]int{1, 2, 3, 4, 5}
sliceFromArray := arr[1:4]  // Contains elements 2, 3, 4
```

Slices have both a length and a capacity:

- **Length**: Number of elements in the slice
- **Capacity**: Maximum number of elements that can be stored without reallocation

## Maps

Maps are key-value pairs:

```go
// String to string map
carInfo := map[string]string{
    "brand": "Ford",
    "model": "Mustang",
    "year":  "1964",
}

// String to int map
cityPopulation := map[string]int{
    "New York": 8419000,
    "Tokyo":    37400000,
    "London":   8982000,
}
```

## Structs

Structs group related data:

```go
type Person struct {
    name   string
    age    int
    job    string
    salary int
}

// Creating and initializing a struct
var employee Person
employee.name = "John"
employee.age = 30
employee.job = "Developer"
employee.salary = 75000

// Alternative initialization
manager := Person{
    name:   "Alice",
    age:    42,
    job:    "Manager",
    salary: 90000,
}
```

## Functions

Functions are defined with the `func` keyword:

```go
// Simple function
func greet() {
    fmt.Println("Hello, world!")
}

// Function with parameters
func greetPerson(name string) {
    fmt.Println("Hello,", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Recursive function
func countdown(n int) {
    if n <= 0 {
        fmt.Println("Liftoff!")
        return
    }
    fmt.Println(n)
    countdown(n - 1)
}
```

## User Input

Go provides several ways to get user input:

```go
// Basic input (space-separated values)
var name string
fmt.Scan(&name)

// Multiple values
var age, height int
fmt.Scan(&age, &height)

// Formatted input
var day, month, year int
fmt.Scanf("%d/%d/%d", &day, &month, &year)

// Read a full line, including spaces
var sentence string
reader := bufio.NewReader(os.Stdin)
sentence, _ = reader.ReadString('\n')
```

## Control Structures

### Loops

```go
// Basic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-like loop
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    // Use break to exit
    break
}

// Iterating over collections
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

### Conditionals

```go
// If statement
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x > 5 {
    fmt.Println("x is between 5 and 10")
} else {
    fmt.Println("x is 5 or less")
}

// Switch statement
day := 3
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Invalid day")
}
```

## Cleaned Example Code

Below is a cleaned version of the original code with better structure, meaningful variable names, and improved organization:

```go
package main

import (
    "fmt"
)

// Person defines the structure for storing person information
type Person struct {
    name   string
    age    int
    job    string
    salary int
}

// displayMessage prints a simple message
func displayMessage() {
    fmt.Println("Function executed successfully!")
}

// greetPerson greets a person with their family name
func greetPerson(firstName string) {
    fmt.Println("Hello", firstName, "Refsnes")
}

// countToTarget counts from start to target, then returns 0
func countToTarget(start int, target int) int {
    if start == target {
        return 0
    }
    fmt.Println(start)
    return countToTarget(start+1, target)
}

func main() {
    // Basic variable declarations
    var name string = "gopher"
    var number int = 10
    var point float32 = 10.24
    var isActive bool = true
    nickname := "coder" // Type inference

    // Constant declaration
    const AppName string = "GoDemo"

    // Arrays
    numbers := [6]int{10, 11, 12, 13, 14, 15}
    moreNumbers := [...]int{4, 5, 6, 7, 8, 9, 10} // Size determined by initializer

    // Maps
    carDetails := map[string]string{
        "brand": "Ford",
        "model": "Mustang",
        "year":  "1964",
    }

    cityIDs := map[string]int{
        "Oslo":      1,
        "Bergen":    2,
        "Trondheim": 3,
        "Stavanger": 4,
    }

    // Slices
    emptyIntSlice := []int{}
    stringSlice := []string{"Go", "Slices", "Are", "Powerful"}
    numberSlice := numbers[2:4] // Slice from array

    // Creating a struct instance
    employee := Person{
        name:   "Hege",
        age:    45,
        job:    "Teacher",
        salary: 6000,
    }

    // Displaying struct information
    fmt.Println("Name:", employee.name)
    fmt.Println("Age:", employee.age)
    fmt.Println("Job:", employee.job)
    fmt.Println("Salary:", employee.salary)

    // Function calls
    displayMessage()
    greetPerson("Liam")
    countToTarget(1, 11)

    // Displaying maps
    fmt.Printf("Car Details: %v\n", carDetails)
    fmt.Printf("City IDs: %v\n", cityIDs)

    // Displaying slices
    fmt.Println("Empty slice length:", len(emptyIntSlice))
    fmt.Println("Empty slice capacity:", cap(emptyIntSlice))
    fmt.Println("Empty slice:", emptyIntSlice)

    fmt.Println("String slice length:", len(stringSlice))
    fmt.Println("String slice capacity:", cap(stringSlice))
    fmt.Println("String slice:", stringSlice)

    fmt.Printf("Number slice = %v\n", numberSlice)
    fmt.Printf("Number slice length = %d\n", len(numberSlice))
    fmt.Printf("Number slice capacity = %d\n", cap(numberSlice))

    // User input
    var firstNum, secondNum int
    fmt.Print("Type two numbers: ")
    fmt.Scanf("%d %d", &firstNum, &secondNum)
    fmt.Println("Your numbers are:", firstNum, "and", secondNum)

    // Displaying arrays
    fmt.Println("First array:", numbers)
    fmt.Println("Second array:", moreNumbers)

    // Displaying variables
    fmt.Println("name:", name)
    fmt.Println("number:", number)
    fmt.Println("point:", point)
    fmt.Println("isActive:", isActive)
    fmt.Println("nickname:", nickname)
    fmt.Println("AppName:", AppName)

    // More user input examples
    var username string
    fmt.Print("Enter your username: ")
    fmt.Scan(&username)

    var age, height int
    fmt.Print("Enter age and height: ")
    fmt.Scan(&age, &height)
    fmt.Printf("Your age is: %d and height is: %d, %s\n", age, height, username)

    // For loop example
    fmt.Println("Counting from 0 to 4:")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Switch statement example
    day := 4
    fmt.Print("Today is: ")
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day")
    }
}
```
