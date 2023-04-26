package main

import (
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type Dump struct {
	ID          int       `json:"id"`
	App         string    `json:"app"`
	Environment string    `json:"environment"`
	Tenants     string    `json:"tenants"`
	DumpDate    time.Time `json:"dump_date"`
	Author      int      `json:"author"`
}

type CreateDumpRequest struct {
	App         string    `json:"app"`
	Environment string    `json:"environment"`
	Tenants     string    `json:"tenants"`
	DumpDate    time.Time `json:"dump_date"`
	Author      int       `json:"author"`
}
