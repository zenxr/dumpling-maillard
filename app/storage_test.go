package main

import (
	"testing"
    "time"
)

func TestInMemoryUserStorage(t *testing.T) {
	s := MemoryUserStorage{}
	user := User{
		ID:   1,
		Name: "test",
	}
	if err := s.CreateUser(&user); err != nil {
		t.Fatalf(`Failed to create user, error: %v`, err)
	}

	foundUser, err := s.GetUser(user.ID)
	if err != nil || foundUser != &user {
		t.Fatalf(`Failed to fetch user, error: %v`, err)
	}

	users, err := s.GetUsers()
	if err != nil || len(users) != 1 {
		t.Fatalf(`Expected 1 user, error: %v`, err)
	}

	user2 := User{
		ID:   2,
		Name: "test2",
	}
	if err := s.CreateUser(&user2); err != nil {
		t.Fatalf(`Failed to create second user, error: %v`, err)
	}
	users, err = s.GetUsers()
	if err != nil || len(users) != 2 {
		t.Fatalf(`Expected 2 user, got %v & error: %v`, users, err)
	}

	if err = s.DeleteUser(1); err != nil {
		t.Fatalf(`Failed to delete user, error: %v`, err)
	}
	users, err = s.GetUsers()
	if err != nil || len(users) != 1 {
		t.Fatalf(`Expected 2 user, got %v & error: %v`, users, err)
	}
}


func TestInMemoryDumpStorage(t *testing.T) {
	s := MemoryDumpStorage{}
    dump := Dump{
        ID: 1,
        App: "app",
        Environment: "test",
        Tenants: "test_tenant, test_tenant2",
        DumpDate: time.Now(),
        Author: 1,
    }
	if err := s.CreateDump(&dump); err != nil {
		t.Fatalf(`Failed to create dump, error: %v`, err)
	}

	foundDump, err := s.GetDump(dump.ID)
	if err != nil || foundDump != &dump {
		t.Fatalf(`Failed to fetch user, error: %v`, err)
	}

	dumps, err := s.GetDumps()
	if err != nil || len(dumps) != 1 {
		t.Fatalf(`Expected 1 dump, error: %v`, err)
	}

	dump2 := Dump{
        ID: 2,
        App: "app2",
        Environment: "test2",
        Tenants: "test_tenant, test_tenant3",
        DumpDate: time.Now(),
        Author: 2,
	}
	if err := s.CreateDump(&dump2); err != nil {
		t.Fatalf(`Failed to create second dump, error: %v`, err)
	}
	dumps, err = s.GetDumps()
	if err != nil || len(dumps) != 2 {
		t.Fatalf(`Expected 2 dump, got %v & error: %v`, dumps, err)
	}

	if err = s.DeleteDump(1); err != nil {
		t.Fatalf(`Failed to delete dump, error: %v`, err)
	}
	dumps, err = s.GetDumps()
	if err != nil || len(dumps) != 1 {
		t.Fatalf(`Expected 2 dump, got %v & error: %v`, dumps, err)
	}
}
