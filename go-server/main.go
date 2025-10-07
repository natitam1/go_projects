package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handles form submission
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n") // fixed typo 'fat' → 'fmt'

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Handles /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" { // must be uppercase 'GET'
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Serve static files from ./static
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting the server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // fixed typo 'Fetal' → 'Fatal'
	}
}
