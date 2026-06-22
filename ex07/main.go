package main

import(
	"fmt"
	"log"
	"net/http"
)

// It accepts an http.ResponseWriter and an *http.Request and returns nothing.
func legacyHandler(w http.ResponseWriter, r *http.Request)  {
	// Redirects the client to /v2
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

// It accepts an http.ResponseWriter and an *http.Request and returns nothing.
func v2Handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Welcome to version 2") // v2Handler writes a welcome message for version 2.
}

func main()  {
	http.HandleFunc("/legacy", legacyHandler)
	http.HandleFunc("/v2", v2Handler)
	fmt.Println("Server running at http://localhost:8080/legacy")
	log.Fatal(http.ListenAndServe(":8080", nil))
}