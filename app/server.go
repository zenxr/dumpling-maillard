package main

import (
	"log"
	"net/http"
    "os"
    "encoding/json"
)

func main() {
	server := ServerFromEnv()
	server.Run()
}

type Server struct {
	address string
	storage Storage
}

func ServerFromEnv() *Server {
	return &Server{
		address: readEnv("address", ":5000"),
		storage: *CreateMemoryStorage(),
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", makeUserHTTPHandleFunc(HandleUser, s.storage.UserStorage))
	mux.HandleFunc("/user/{id}", makeUserHTTPHandleFunc(HandleUserByID, s.storage.UserStorage))
	mux.HandleFunc("/dump", makeDumpHTTPHandleFunc(HandleDump, s.storage.DumpStorage))
	mux.HandleFunc("/dump/{id}", makeDumpHTTPHandleFunc(HandleDumpByID, s.storage.DumpStorage))

	log.Println("Running on port: ", s.address)
	http.ListenAndServe(s.address, mux)
}

func readEnv(key string, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string `json:"error"`
}

type UserRequestHandlerFunction func(s UserStorage, w http.ResponseWriter, r *http.Request) error

func makeUserHTTPHandleFunc(handler UserRequestHandlerFunction, storage UserStorage) http.HandlerFunc {
	// injects dependencies, reduces error handling
	// could be generic, if I only didn't hate exposing a `GetId` to make
	// an interface
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(storage, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type DumpRequestHandlerFunction func(s DumpStorage, w http.ResponseWriter, r *http.Request) error

func makeDumpHTTPHandleFunc(handler DumpRequestHandlerFunction, storage DumpStorage) http.HandlerFunc {
	// injects dependencies, reduces error handling
	// could be generic, if I only didn't hate exposing a `GetId` to make
	// an interface
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(storage, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
