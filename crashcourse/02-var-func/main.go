package main

import "fmt"

var globalVar = "This is the global variable"
// var GlobalVar = "The variable is Public"

func main(){
	// declaration of a variable
	var awesomeString string
	var coolInteger int
	
	// assignment of variable
	awesomeString = "Very Awesome!"
	coolInteger = 43

	// declaration + assignment (walrus operator)
	greatFloat := 43.3

	fmt.Println("awesomeString is:", awesomeString)
	fmt.Println("coolInteger is:", coolInteger)
	fmt.Println("greatFloat is:", greatFloat)
	fmt.Println("globalVar", globalVar)

	// firstName := getFirstname()
	// fmt.Println("FirstName is", firstName)
	fmt.Println("FirstName is", getFirstname())
	
	john, age := getNames()
	fmt.Println("name:", john)
	fmt.Println("age:", age)

	incrementedNumber := incrementByOne(66)
	fmt.Println("increment by one:", incrementedNumber)

	twoFloats := addTwoFloats(45.3, 54.5)
	fmt.Println("Result is:", twoFloats)

	subThis := SubtractNumbers(45, 24)
	fmt.Println("Result is:", subThis)
}

func getFirstname() string {
	return "fsec"
}

func getNames() (string, int){
	return "John", 34
}

func incrementByOne(number int) int{
	return number + 1
}

func addTwoFloats(numberOne, numbertwo float64) float64{
	return numberOne + numbertwo
}

// a capital in front of function or variable, it means public. You can use it in others file.
func SubtractNumbers(numberOne, numbertwo int) int{
	return numberOne - numbertwo
}