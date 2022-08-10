package main

import (
	"errors"
	"log"
)

var (
	errBadThing   = errors.New("something bad")
	errWorseThing = errors.New("worse thing")
)

func main() {
	err := doSomethingBad()
	_ = doSomethingWorse()

	if err != nil {
		switch err {
		case errBadThing:
			log.Printf("Uh oh: %s", err)
		case errWorseThing:
			panic("Abandon ship!")
		}
	}
}

func doSomethingBad() error {
	return errBadThing
}

func doSomethingWorse() error {
	return errWorseThing
}
