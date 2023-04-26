package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HandleUser(s UserStorage, w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getUsers(s, w, r)
	}
	if r.Method == "POST" {
		return createUser(s, w, r)
	}
	return fmt.Errorf("Invalid method: %s", r.Method)
}

func HandleUserByID(s UserStorage, w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getUserById(s, w, r)
	}
	if r.Method == "DELETE" {
		return deleteUserById(s, w, r)
	}
	return fmt.Errorf("Invalid method: %s", r.Method)
}

func getUsers(s UserStorage, w http.ResponseWriter, r *http.Request) error {
	users, err := s.GetUsers()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, users)
}

func createUser(s UserStorage, w http.ResponseWriter, r *http.Request) error {
	req := new(CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	user := User{Name: req.Name}
	s.CreateUser(&user)
	return writeJSON(w, http.StatusCreated, user)
}

func getUserById(s UserStorage, w http.ResponseWriter, r *http.Request) error {
	id_param := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(id_param)
	if err != nil {
		return errors.New("Expected id of type integer")
	}

	user, err := s.GetUser(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, user)
}

func deleteUserById(s UserStorage, w http.ResponseWriter, r *http.Request) error {
	id_param := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(id_param)
	if err != nil {
		return errors.New("Expected id of type integer")
	}
	if err := s.DeleteUser(id); err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}
