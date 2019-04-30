package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	httpHost := os.Getenv("PRODUCER_HOST")
	httpPort := os.Getenv("PRODUCER_PORT")

	createHost := os.Getenv("CREATOR_HOST")
	createPort := os.Getenv("CREATOR_PORT")

	http.HandleFunc("/input", func(w http.ResponseWriter,
		r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			resp, err := http.Get(fmt.Sprintf("http://%v:%v/sayHello?name=%v", createHost, createPort, name))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
			} else {
				byteToWrite, _ := ioutil.ReadAll(resp.Body)
				w.Write(byteToWrite)
			}
			defer resp.Body.Close()
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("name parameter not found"))
		}

	})

	http.ListenAndServe(fmt.Sprintf("%v:%v", httpHost, httpPort), nil)
	log.Printf("running on: %v:%v\n", httpHost, httpPort)
}
