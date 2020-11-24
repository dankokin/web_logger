package models

import (
	"time"
)

type Config struct {
	ConnectionString string `json:"connection_string"`
	ServerPort int `json:"port"`
}

type WAFMessage struct {
	RegisteredAt time.Time `json:"registred_at"`
	Message      string `json:"message"`
	Target       string `json:"target"`
}

