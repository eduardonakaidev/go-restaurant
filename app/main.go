package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello, world!!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", MainHandler)
	fmt.Println("Listening on port 5050...")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		// Handle the error here
		log.Fatal(err)
	}
}
