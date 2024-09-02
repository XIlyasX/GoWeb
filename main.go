package main

import (
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server!"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/message", messageHandler)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
