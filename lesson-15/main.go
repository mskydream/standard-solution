package main

// *********      code old version     **********

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "I'm root.")
// }

// type proverb struct {
// 	id    int
// 	value string
// }

// type proverbsHandler struct {
// 	proverbs []proverb
// }

// func newProverbsHandler() *proverbsHandler {
// 	return &proverbsHandler{
// 		proverbs: []proverb{
// 			proverb{id: 1, value: "Don't panic"},
// 			proverb{id: 2, value: "Concurrency is not parallelism."},
// 			proverb{id: 3, value: "Documentation is for users."},
// 			proverb{id: 4, value: "The bigger the interface, the weaker the abstraaction."},
// 			proverb{id: 5, value: "Make the zero value useful"},
// 		},
// 	}
// }

// func (ph *proverbsHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Path[len("/proverbs/"):])
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	p, err := ph.lookup(id)
// 	if err == errUnknownProverb {
// 		http.Error(w, errUnknownProverb.Error(), http.StatusNotFound)
// 	}

// 	fmt.Fprintln(w, p.value)
// }

// type greetHandler struct{}

// func (ph *greetHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello Gopher!")
// }

// func main() {
// 	// gh := &greetHandler{}
// 	// http.Handle("/greet/", gh)

// 	ph := newProverbsHandler()
// 	http.Handle("/proverbs/", ph)

// 	http.HandleFunc("/", handler)

// 	log.Println("Starting server...")
// 	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
// }
