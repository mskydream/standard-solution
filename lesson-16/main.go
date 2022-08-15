package main

// old version code

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync/atomic"
// 	"time"
// )

// var requestHandled uint64

// func greetHandler(w http.ResponseWriter, r *http.Request) {
// 	atomic.AddUint64(&requestHandled, 1)

// 	log.Printf("STAT %s %q\n", r.Method, r.URL.String())
// 	t := time.Now()

// 	fmt.Fprintln(w, "Hello")
// 	log.Println("GREET")

// 	log.Printf("END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
// }

// type statsHandler struct{}

// func (sh *statsHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("STAT %s %q\n", r.Method, r.URL.String())
// 	t := time.Now()

// 	fmt.Fprintf(w, "Request Handled: %d\n", atomic.LoadInt64(&requestHandled))
// 	log.Println("STATS PROVIDED")

// 	log.Printf("END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
// }

// func main() {
// 	http.HandleFunc("/greet", greetHandler)

// 	sh := &statsHandler{}
// 	http.Handle("/stats", sh)

// 	log.Println("Starting server...")
// 	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
// }
