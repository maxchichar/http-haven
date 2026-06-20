package main

import(
	"log"
	"fmt"
	"net/http"
)

func agentHandler(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")
	if agent == "" {
		agent = "Unknown device" // Fall-back response if device not recognised
	}

	fmt.Fprintf(w, "You are visiting us using: %s", agent) // Prints the message with the device info
}

// Server setup
func main()  {
	http.HandleFunc("/agent", agentHandler)
	fmt.Println("Server running at http://localhost:8080/agent")
	log.Fatal(http.ListenAndServe(":8080", nil))
}