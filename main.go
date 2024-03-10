package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not fund", http.StatusNotFound)

		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supportted", http.StatusNotFound)
	}

	fmt.Fprint(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return

	}
	fmt.Fprintf(w, "Post request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Address = %s", address)

}

func formHander1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "from.html")
}

func main() {

	// Create a file server
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form1", formHander1)

	fmt.Println("Starting server at port: 8081")

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}
}
