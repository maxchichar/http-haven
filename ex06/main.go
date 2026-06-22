package main

import(
	"fmt"
	"log"
	"net/http"
)

// dashboardHandler handles requests to the dashboard endpoint.
//
func dashboardHandler(w http.ResponseWriter, r *http.Request)  { // w: http.ResponseWriter used to write the response.
	taken := r.Header.Get("X-API-Key") // taken gets the X-API-Key header value from request r.
	if taken != "secret123" { // It checks for the X-API-Key header; if the value equals "secret123",
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome") // it writes a welcome message, otherwise it returns HTTP 401 Unauthorized.
}

func main()  {
	http.HandleFunc("/dashboard", dashboardHandler) // HandleFunc registers dashboardHandler to serve requests at the /dashboard path.
	fmt.Println("Server running at http://localhost:8080/dashboard")
	log.Fatal(http.ListenAndServe(":8080", nil))
}