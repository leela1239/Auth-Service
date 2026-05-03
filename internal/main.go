package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	registerRoutes(mux)

	fmt.Println("Server running on :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
