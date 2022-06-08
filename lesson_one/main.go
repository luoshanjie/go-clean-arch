package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "status response")
}

func main() {
	http.HandleFunc("/api/v1/user/add", handleAddUser)
	log.Fatal(http.ListenAndServe(":10086", nil))
}
