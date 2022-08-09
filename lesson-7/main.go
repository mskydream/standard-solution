package main

import (
	"fmt"
	"time"
)

// lesson with terminal
var version string

func main() {
	fmt.Println("Hello Gophers!")
	fmt.Printf("Current time: %v\n", time.Now())
	fmt.Printf("Version: %s\n", version)
}
