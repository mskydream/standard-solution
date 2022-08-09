package mypkg

import "fmt"

func SayHello(name string) {
	fmt.Printf("Hello %s!", name)
}

type hidden struct{}
