package main

import (
	"encoding/json"
	"fmt"
)

// Book is ...
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author is ...
type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

// SensorReading is ...
type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func marshallingJSON() {
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	byteArrayIndent, errInden := json.MarshalIndent(book, "", "  ")
	if errInden != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArrayIndent))
}

func unmarshallingJSON() {
	unmarshallingJSONIntoStructuredData()
	unmarshallingJSONIntoUnsructuredData()
}

func unmarshallingJSONIntoStructuredData() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}

func unmarshallingJSONIntoUnsructuredData() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}

func main() {
	marshallingJSON()
	unmarshallingJSON()
}
