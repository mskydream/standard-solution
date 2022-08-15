package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path"
)

type Proverb struct {
	ID       int
	Text     string
	reviewed bool
}

func (p Proverb) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	_, err := fmt.Fprintf(&b, "ID=%d Text=%q reviewed=%t\n", p.ID, p.Text, p.reviewed)
	return b.Bytes(), err
}

func main() {
	proverbs := []Proverb{
		{ID: 1, Text: "Don't panic.", reviewed: true},
		{ID: 2, Text: "Concurrency is not parallelism.", reviewed: true},
		{ID: 3, Text: "Documentation is for users.", reviewed: true},
		{ID: 4, Text: "The bigger the interface, the weaker the abstraction.", reviewed: true},
		{ID: 5, Text: "Make the zero value useful.", reviewed: true},
	}

	filename := path.Join("..", "proverbs.gob")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	if err := enc.Encode(proverbs); err != nil {
		log.Fatalln(err)
	}

	log.Println("Done.")
}
