package main

// import the 2 modules we need
import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(filename string) {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile(filename)
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Print(string(data))
}

func writeFile(filename string) {
	mydata := []byte("all my data I want to write to a file\n")

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile(filename, mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

func appendFile(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}
}

func removeFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Sprintln("File %s successfully deleted.", filename)
}

func main() {
	readFile("localfile.data")
	writeFile("myfile.data")
	appendFile("myfile.data")
	readFile("myfile.data")
	removeFile("myfile.data")
}
