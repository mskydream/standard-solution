package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type Proverb struct {
	ID   int
	Text string
}

func main() {
	filename := path.Join("..", "proverbs.json")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var proverbs []Proverb

	dec := json.NewDecoder(file)
	if err := dec.Decode(&proverbs); err != nil {
		log.Fatalln(err)
	}

	for _, p := range proverbs {
		log.Printf("%#v", p)
	}
}
