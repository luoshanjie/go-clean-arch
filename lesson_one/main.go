package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cache map[string]User

func handleAddUser(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "status response")
}

func main() {
	http.HandleFunc("/api/v1/user/add", handleAddUser)
	log.Fatal(http.ListenAndServe(":80", nil))
}
