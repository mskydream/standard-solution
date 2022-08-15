package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handler greeting request")

	completeAfter := time.After(5 * time.Second)

	for {
		select {
		case <-completeAfter:
			fmt.Fprintln(w, "Hello Gopher")
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Greeting are hard. Thinking...")
		}
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
