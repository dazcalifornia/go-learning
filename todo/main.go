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
	numberSlice := numbers[0:4] // Slice from array

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
