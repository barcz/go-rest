package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type resp struct {
	Message string `json:"message"`
}

func sayHello(name string) ([]byte, error) {
	response := &resp{
		Message: fmt.Sprintf("Hello %v", name),
	}
	return json.Marshal(response)
}

func main() {

	host := os.Getenv("CREATOR_HOST")
	port := os.Getenv("CREATOR_PORT")

	http.HandleFunc("/sayHello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("name url parameter not found"))
		} else {
			hello, _ := sayHello(name)
			log.Printf("saying hello to %v", name)
			w.Write(hello)
		}
	})

	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil)
}
