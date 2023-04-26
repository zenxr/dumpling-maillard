package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HandleDump(s DumpStorage, w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getDumps(s, w, r)
	}
	if r.Method == "POST" {
		return createDump(s, w, r)
	}
	return fmt.Errorf("Invalid method: %s", r.Method)
}

func HandleDumpByID(s DumpStorage, w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return getDumpById(s, w, r)
	}
	if r.Method == "DELETE" {
		return deleteDumpById(s, w, r)
	}
	return fmt.Errorf("Invalid method: %s", r.Method)
}

func getDumps(s DumpStorage, w http.ResponseWriter, r *http.Request) error {
	dumps, err := s.GetDumps()
	if err != nil {
        return err
	}
	return writeJSON(w, http.StatusOK, dumps)
}

func createDump(s DumpStorage, w http.ResponseWriter, r *http.Request) error {
	req := new(CreateDumpRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
    dump := Dump{
        App: req.App,
        Environment: req.Environment,
        Tenants: req.Tenants,
        DumpDate: req.DumpDate,
        Author: req.Author, // todo: fk relationship
    }
	s.CreateDump(&dump)
	return writeJSON(w, http.StatusCreated, dump)
}

func getDumpById(s DumpStorage, w http.ResponseWriter, r *http.Request) error {
	id_param := strings.TrimPrefix(r.URL.Path, "/dump/")
	id, err := strconv.Atoi(id_param)
	if err != nil {
        return errors.New("Expected id of type integer")
	}
	dump, err := s.GetDump(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, dump)
}

func deleteDumpById(s DumpStorage, w http.ResponseWriter, r *http.Request) error {
	id_param := strings.TrimPrefix(r.URL.Path, "/dump/")
	id, err := strconv.Atoi(id_param)
	if err != nil {
        return errors.New("Expected id of type integer")
	}
	if err := s.DeleteDump(id); err != nil {
		return err
	}
    return writeJSON(w, http.StatusNoContent, nil)
}
