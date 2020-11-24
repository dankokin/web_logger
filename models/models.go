package models

import (
	"time"
)

// DataBase options
type DataBase struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"ssl_mode"`
}

// Server options
type Server struct {
	Port int `yaml:"port"`
}

// Structure for
type WAFMessage struct {
	RegisteredAt time.Time `json:"registred_at"`
	Message      string `json:"message"`
	Target       string `json:"target"`
}

// Convenient structure for launching the service
type Config struct {
	Server   `yaml:"server"`
	DataBase `yaml:"data_base"`
}
