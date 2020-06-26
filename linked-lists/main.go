package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Go linked Lists Tutorial")
	myList := list.New()
	myList.PushBack(1)
	myList.PushFront(2)

	fmt.Println("Inserting elements")
	for element := myList.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

	fmt.Println("Removing elements")
	for element := myList.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		if element.Value != 1 {
			myList.Remove(element)
		}
	}

	fmt.Println("Removed elements")
	for element := myList.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}
}
