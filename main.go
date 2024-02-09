package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /\n")
	io.WriteString(w, "Hi there!!\n")
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /data\n")
    fetchAPI()
	io.WriteString(w, "Hi there!!\n")
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/data", getData)
    log.Println("Starting server on :8000")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}
