package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    message := Message{Text: "Hello, World!"}
    json.NewEncoder(w).Encode(message)
}