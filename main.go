package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Printf(w, "Name: %v, Address: %v", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Invalid URL", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello %s", r.URL.Path)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 5050\n")
	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Fatal(err)
	}
}
