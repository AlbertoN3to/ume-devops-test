package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got ping")
	fmt.Fprintf(w, "Pong")
}

func main() {
	http.HandleFunc("/api/v1/pong", pong)

	fmt.Printf("running server on port 9090\n")

	http.ListenAndServe(":9090", nil)
}
