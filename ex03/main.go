package main

import(
	"io"
	"log"
	"fmt"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		fmt.Fprint(w, "Send a POST request with text to count words")
	}else if r.Method == "POST"{
		body, err := io.ReadAll(r.Body)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%d", len(body)) // calculate the length of the 
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/count", countHandler)
	fmt.Println("Server running at http://localhost:8080/count")
	log.Fatal(http.ListenAndServe(":8080", nil))
}