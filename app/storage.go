package main

import (
	"errors"
	"fmt"
)

type UserStorage interface {
	CreateUser(*User) error
	DeleteUser(id int) error
	GetUser(id int) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(*User) error
}

type DumpStorage interface {
	CreateDump(*Dump) error
	DeleteDump(id int) error
    GetDump(id int) (*Dump, error)
	GetDumps() ([]*Dump, error)
	UpdateDump(*Dump) error
}

type Storage struct {
	UserStorage UserStorage
	DumpStorage DumpStorage
}

func CreateMemoryStorage() *Storage {
	return &Storage{&MemoryUserStorage{}, &MemoryDumpStorage{}}
}

type MemoryUserStorage struct {
	users []*User
}

func (s *MemoryUserStorage) CreateUser(u *User) error {
	s.users = append(s.users, u)
	return nil
}

func (s *MemoryUserStorage) DeleteUser(id int) error {
	n := 0
	for _, user := range s.users {
		if user.ID != id {
			s.users[n] = user
			n++
		}
	}

	if n == len(s.users) {
		return errors.New("ID not found")
	}

	s.users = s.users[:n]
	return nil
}

func (s *MemoryUserStorage) GetUser(id int) (*User, error) {
    for _, user := range s.users {
        if user.ID == id {
            return user, nil
        }
    }
    return nil, errors.New(fmt.Sprintf("ID %d not found", id))
}

func (s *MemoryUserStorage) GetUsers() ([]*User, error) {
	return s.users, nil
}

func (s *MemoryUserStorage) UpdateUser(u *User) error {
	for idx, user := range s.users {
		if user.ID == u.ID {
			s.users[idx] = u
		}
		return nil
	}
	return errors.New("User not found")
}

type MemoryDumpStorage struct {
	dumps []*Dump
}

func (s *MemoryDumpStorage) CreateDump(d *Dump) error {
	s.dumps = append(s.dumps, d)
	return nil
}

func (s *MemoryDumpStorage) DeleteDump(id int) error {
	n := 0
	for _, dump := range s.dumps {
		if dump.ID != id {
			s.dumps[n] = dump
			n++
		}
	}

	if n == len(s.dumps) {
		return errors.New("ID not found")
	}

	s.dumps = s.dumps[:n]
	return nil
}

func (s *MemoryDumpStorage) GetDump(id int) (*Dump, error) {
    for _, dump := range s.dumps {
        if dump.ID == id {
            return dump, nil
        }
    }
    return nil, errors.New(fmt.Sprintf("ID %d not found", id))
}

func (s *MemoryDumpStorage) GetDumps() ([]*Dump, error) {
	return s.dumps, nil
}

func (s *MemoryDumpStorage) UpdateDump(d *Dump) error {
	for idx, dump := range s.dumps {
		if dump.ID == d.ID {
			s.dumps[idx] = d
		}
		return nil
	}
	return errors.New("Dump not found")
}
