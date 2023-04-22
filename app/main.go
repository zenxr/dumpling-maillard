
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users []User

func main() {
    // Define routes
    http.HandleFunc("/users", getUsers)
    http.HandleFunc("/users/add", addUser)

    // Start server
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    // Return list of users as JSON
    json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
    // Parse request body to get user data
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Add user to list
    users = append(users, user)

    // Return added user as JSON
    json.NewEncoder(w).Encode(user)
}
