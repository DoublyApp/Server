package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex

func get_users(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUsers(w, r))
	mu.Unlock()
}

func get_user_by_id(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUserByID(w, r))
	mu.Unlock()
}

func main() {
    http.HandleFunc("/get_users/", get_users)
    http.HandleFunc("/get_user_by_id/", get_user_by_id)
    log.Fatal(http.ListenAndServe(":80", nil))
}
