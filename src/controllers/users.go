package controllers

import "net/http"

// CreateUser inserts a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating User!"))
}

// SearchUsers search all users saved in the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating User!"))
}

// SearchUser searches for a user saved in the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating User!"))
}

// Update User changes a user's information in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating User!"))
}

// DeleteUser deletes a user's information in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating User!"))
}