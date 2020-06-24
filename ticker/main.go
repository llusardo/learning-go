package main

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Tock")
	}
}

func main() {
	fmt.Println("Go Tickers Tutorial")

	go backgroundTask()

	select {}
}
