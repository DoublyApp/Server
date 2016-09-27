package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex

// users.go
func get_users(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUsers(w, r))
	mu.Unlock()
}

// users.go
func get_user_by_id(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUserByID(w, r))
	mu.Unlock()
}

// interests.go
func get_interests(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetInterests(w, r))
	mu.Unlock()
}

// interests.go
func get_users_interests(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUsersInterests(w, r))
	mu.Unlock()
}

// interests.go
func insert_interest(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, InsertInterest(w, r))
	mu.Unlock()
}

// interests.go
func remove_interest(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, RemoveInterest(w, r))
	mu.Unlock()
}

// interests.go
// /get_interests_users/<searching user's ID>/<interestID>/<start index>/<end index>
func get_interests_users(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, GetInterestsUsers(w, r))
	mu.Unlock()
}

// GPS.go
// /update_gps/<userID>/<latitude>/<longitude>
func update_gps(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, UpdateGPS(w, r))
	mu.Unlock()
}

// messages.go
// /insert_message/<senderID>/<receiverID>/<messageText>
func insert_message(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, InsertMessage(w, r))
	mu.Unlock()
}

// returns userID to client
// /insert_user/<user name>/<email>/<password>/<DOB>/<gender>
func insert_user(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, InsertUser(w, r))
	mu.Unlock()
}

func main() {
	http.HandleFunc("/get_users/", get_users)
	http.HandleFunc("/get_user_by_id/", get_user_by_id)
	http.HandleFunc("/insert_user/", insert_user)
	http.HandleFunc("/get_interests/", get_interests)
	http.HandleFunc("/get_users_interests/", get_users_interests)
	http.HandleFunc("/insert_interest/", insert_interest)
	http.HandleFunc("/remove_interest/", remove_interest)
	http.HandleFunc("/get_interests_users/", get_interests_users)
	http.HandleFunc("/update_gps/", update_gps)
	http.HandleFunc("/insert_message/", insert_message)
	log.Fatal(http.ListenAndServe(":80", nil))
}
