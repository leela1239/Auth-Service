package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/view", RequirePermission(handleAuth, Permission("read")))

	http.HandleFunc("/edit", RequirePermission(handleAuth, Permission("write")))

	http.HandleFunc("/delete", RequirePermission(handleAuth, Permission("delete")))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}

	fmt.Println("Server startted on :3000...")

}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User has permission to access this resource."))
}

//  curl.exe -H "X-User-ID: User1" http://localhost:3000/view
