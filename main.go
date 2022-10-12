package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func helloPage(w http.ResponseWriter, r *http.Request) {
	var user User = User{FirstName: "Gigel", LastName: "John", Age: 32}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// middleware checks http requests have HTTP Header "Content-Type" set to "application/json"
func jsonHeaderVerifier(endpoint http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		log.Print("Executing middleware")
		headers := r.Header
		val, exists := headers["Content-Type"]
		if exists {
			if val[0] == "application/json" {
				log.Print("Content-Type key header is present with value", val)
				endpoint.ServeHTTP(w, r)
			} else {
				log.Print("Please set HTTP Header Content-Type:application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not Authorised!\nPlease set HTTP header Content-Type:application/json"))
			}
		} else {
			log.Print("Please set HTTP Header Content-Type:application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorised!\nPlease set HTTP header Content-Type:application/json"))
		}
	})
}

func handleRequests() {
	mux := http.NewServeMux()
	mux.Handle("/", jsonHeaderVerifier(helloPage))
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func main() {
	handleRequests()
}
