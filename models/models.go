package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

type Config struct {
	ConnectionString string `json:"connection_string"`
	ServerPort       int    `json:"port"`
}

type WAFMessage struct {
	TargetDomain              string    `json:"target"`
	RequestURI                string    `json:"request_uri"`
	StatusCode                int       `json:"status_code"`
	RequestRulesCheckElapsed  int64     `json:"request_rules_check_elapsed"`
	ResponseRulesCheckElapsed int64     `json:"response_rules_check_elapsed"`
	HTTPElapsed               int64     `json:"http_elapsed"`
	RequestSize               int64     `json:"request_size"`
	ResponseSize              int64     `json:"response_size"`
	RegisteredAt              time.Time
}

func (msg WAFMessage) Validate() error {
	return validation.ValidateStruct(&msg,
		validation.Field(&msg.TargetDomain, is.URL),
		validation.Field(&msg.RequestURI, is.RequestURI))
}

type ServerResponse struct {
	Logs []WAFMessage `json:"logs"`
}
