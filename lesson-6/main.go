package main

import (
	"fmt"
	"regexp"
)

func main() {
	match()
}

func match() {
	p := "(foo|faux)bar"

	s1 := "fallbar"
	match, err := regexp.MatchString(p, s1)
	fmt.Printf("MatchString - %v, %v\n", match, err)

	s2 := "foobar"
	match, err = regexp.MatchString(p, s2)
	fmt.Printf("MatchString - %v, %v\n", match, err)

}
