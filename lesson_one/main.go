package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cache map[string]

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	fmt.Printf("%+v", user)

	fmt.Fprint(w, "status response")
}

func main() {
	http.HandleFunc("/api/v1/user/add", handleAddUser)
	log.Fatal(http.ListenAndServe(":80", nil))
}
