package main

import (
	// import our encoding/json package
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func openJSONFile(filename string) (*os.File, error) {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %v", filename)
	// defer the closing of our jsonFile so that we can parse it later on
	//defer jsonFile.Close()
	return jsonFile, err
}

func processJSONFile(byteValue []byte) {
	var users Users

	json.Unmarshal(byteValue, &users)

	fmt.Printf("Entering %v", len(users.Users))
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}

func processUnstructuredJSONFile(byteValue []byte) {
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println("------------------------")
	fmt.Println(result["users"])
}

func main() {
	jsonFile, err := openJSONFile("users.json")
	if err != nil {
		fmt.Println(err)
	} else {
		// read our opened jsonFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)
		processJSONFile(byteValue)
		processUnstructuredJSONFile(byteValue)
	}
	jsonFile.Close()
}
