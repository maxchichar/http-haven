package main

import(
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong") // Displays pong on the server
}

func main() {
	http.HandleFunc("/ping", handler)
	fmt.Println("Server running at http://localhost:8080/ping")
	log.Fatal(http.ListenAndServe(":8080", nil)) // the serve should run at :8080
}