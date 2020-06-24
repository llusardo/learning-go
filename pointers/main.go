package main

import "fmt"

func variablePassedByCopyTestFunc(a int) {
	a += 3
	fmt.Println(a)
}

func variablePassedByCopy() {
	fmt.Println("START variablePassedByCopy")
	a := 2
	variablePassedByCopyTestFunc(a)
	fmt.Println(a) // prints out 2
	fmt.Println("END variablePassedByCopy")
}

func variablePassedByValueTestFunc(a *int) {
	*a += 3
	fmt.Println(*a)
}

func variablePassedByValue() {
	fmt.Println("START variablePassedByPointer")
	a := 2
	variablePassedByValueTestFunc(&a)
	fmt.Println(a) // prints out 2
	fmt.Println("END variablePassedByPointer")
}

func definingPointers() {
	var age *int
	age = new(int)
	*age = 26

	fmt.Println(*age)
	//Address of variable
	fmt.Printf("Memory Address of age variable %s", &age)
}

func main() {
	variablePassedByCopy()
	variablePassedByValue()
	definingPointers()
}
