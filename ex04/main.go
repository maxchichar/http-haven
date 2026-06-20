package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	op := r.URL.Query().Get("op")

	a := r.URL.Query().Get("a")
	A, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b := r.URL.Query().Get("b")
	B, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result int

	switch op {
	case "add":
		result = A + B
	case "subtract":
		result = A - B
	case "multiply":
		result = A * B
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}

func main()  {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server running at http://localhost:8080/calculate")
	log.Fatal(http.ListenAndServe(":8080", nil))
}