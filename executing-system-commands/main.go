package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func executeCommand(command string, arg ...string) {
	out, err := exec.Command(command, arg[0:]...).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	// []byte convert to a string
	output := string(out[:])
	fmt.Printf("Command %s Successfully Executed with result: %v", command, output)
}

func execute() {
	executeCommand("ls", "-ltr")
	executeCommand("pwd")
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
