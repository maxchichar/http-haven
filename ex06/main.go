package main

import(
	"fmt"
	"log"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request)  {
	taken := r.Header.Get("X-API-Key")
	if taken != "secret123" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome")
}

func main()  {
	http.HandleFunc("/dashboard", dashboardHandler)
	fmt.Println("Server running at http://localhost:8080/dashboard")
	log.Fatal(http.ListenAndServe(":8080", nil))
}