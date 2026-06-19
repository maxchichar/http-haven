package main

import(
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Method check
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed) // To return 404 error if page not found
		return
	}

	name := r.URL.Query().Get("name")
	if name == ""{ 
		name = "Guest" // created a fallback logic for name
	}

	fmt.Fprintf(w, "Hello, %s!", name) 
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Location at which the server is hosted
}