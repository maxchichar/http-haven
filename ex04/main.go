package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Peformed method check that wil return HTTP error if page not found
	if r.Method != "GET"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	op := r.URL.Query().Get("op")
	
	// Conversion of string to integer using the Atoi in other to perform Arithemetic calculations
	a := r.URL.Query().Get("a")
	A, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // This return HTTP error 400
		return
	}

	b := r.URL.Query().Get("b")
	B, err := strconv.Atoi(b) 
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result int // This stores the output data

	// Arithemetic Formula
	switch op {
	case "add":
		result = A + B
	case "subtract":
		result = A - B
	case "multiply":
		result = A * B
	default:
		w.WriteHeader(http.StatusBadRequest) // Returns error if the operator is unknown
		return
	}

	fmt.Fprintf(w, "Result: %d", result) // What the user see as the server output
}

func main()  {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server running at http://localhost:8080/calculate") // The URL to perform you're calculations
	log.Fatal(http.ListenAndServe(":8080", nil)) // Server Address where it is hosted at :8080
}