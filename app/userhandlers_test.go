package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	s := MemoryUserStorage{}
	handler := http.HandlerFunc(makeUserHTTPHandleFunc(HandleUser, &s))

	addStubUsers(&s)

	rr := simulateRequest(handler, "GET", "/user", nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"name":"test"},{"id":2,"name":"test1"}]`
	require.JSONEq(t, expected, rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	s := MemoryUserStorage{}
	handler := http.HandlerFunc(makeUserHTTPHandleFunc(HandleUser, &s))

	requestBody := `{"name": "test"}`

	rr := simulateRequest(handler, "POST", "/user", &requestBody)

	expectedStatus := http.StatusCreated
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}
	expected := `{"id":0, "name":"test"}`
	require.JSONEq(t, expected, rr.Body.String())

	// user was persisted to store
	if users, _ := s.GetUsers(); len(users) != 1 {
		t.Errorf("expected single %d users, got %d",
			1, len(users))
	}
}

func TestGetUserById(t *testing.T) {
	s := MemoryUserStorage{}
	handler := http.HandlerFunc(makeUserHTTPHandleFunc(HandleUserByID, &s))

	addStubUsers(&s)
	rr := simulateRequest(handler, "GET", "/user/2", nil)

	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	expected := `{"id":2, "name":"test1"}`
	require.JSONEq(t, expected, rr.Body.String())
}

func TestDeleteUser(t *testing.T) {
	s := MemoryUserStorage{}
	handler := http.HandlerFunc(makeUserHTTPHandleFunc(HandleUserByID, &s))

	addStubUsers(&s)
	rr := simulateRequest(handler, "DELETE", "/user/1", nil)

	expectedStatus := http.StatusNoContent
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	// one user remaining
	if users, _ := s.GetUsers(); len(users) != 1 {
		t.Errorf("expected single %d users, got %d",
			1, len(users))
	}
}
func simulateRequest(handler http.Handler, method string, route string, body *string) *httptest.ResponseRecorder {
	var req *http.Request
	if body == nil {
		req = httptest.NewRequest(method, route, nil)
	} else {
		req = httptest.NewRequest(method, route, strings.NewReader(*body))
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	return rr
}

func addStubUsers(s UserStorage) {
	users := []*User{
		{ID: 1, Name: "test"},
		{ID: 2, Name: "test1"},
	}
	for _, user := range users {
		s.CreateUser(user)
	}
}
