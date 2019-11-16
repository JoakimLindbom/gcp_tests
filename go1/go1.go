package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	//initPubSub()

	r := mux.NewRouter()

	r.HandleFunc("/", baseHandler)
	//r.HandleFunc("/", baseHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	r.Handle("/",r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port),r))

}



func baseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World! 123")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World! 123")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
