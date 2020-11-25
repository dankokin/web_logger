package services

import (
	"strings"
	"time"

	"github.com/dankokin/web_logger/models"
)

type DataStore interface {
	// Function that gets all logs from database by filters
	GetAllLogs(map[string]interface{}) (models.ServerResponse, error)

	// Simple function that saves WAFMessage to database
	SaveMessage(models.WAFMessage) error
}

func (db *DB)GetAllLogs(filters map[string]interface{}) (models.ServerResponse, error) {
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("SELECT target, request_uri, status_code, request_rules_check_elapsed, " +
		"response_rules_check_elapsed, http_elapsed, request_size, response_size, registered_at FROM Messages where registered_at >= $1")

	args := make([]interface{}, 0, 2)

	interval := filters["interval"].(int)
	beginInterval := time.Now().Add(-24 * time.Hour * time.Duration(interval))

	args = append(args, beginInterval)

	if target, exists := filters["target"]; exists {
		args = append(args, target)
		queryBuilder.WriteString(" and target = $2")
	}

	rows, err := db.Query(queryBuilder.String(), args...)
	if err != nil {
		return models.ServerResponse{}, err
	}

	messages := make([]models.WAFMessage, 0, 256)
	for rows.Next() {
		var msg models.WAFMessage
		err = rows.Scan(
			&msg.TargetDomain,
			&msg.RequestURI,
			&msg.StatusCode,
			&msg.RequestRulesCheckElapsed,
			&msg.ResponseRulesCheckElapsed,
			&msg.HTTPElapsed,
			&msg.RequestSize,
			&msg.ResponseSize,
			&msg.RegisteredAt)

		if err != nil {
			return models.ServerResponse{}, err
		}

		messages = append(messages, msg)
	}

	return models.ServerResponse{Logs: messages}, nil
}

func (db *DB)SaveMessage(message models.WAFMessage) error {
	_, err := db.Exec("INSERT INTO Messages (target, request_uri, status_code, request_rules_check_elapsed," +
		" response_rules_check_elapsed, http_elapsed, request_size, response_size, registered_at) VALUES " +
		"($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		message.TargetDomain, message.RequestURI,
		message.StatusCode, message.RequestRulesCheckElapsed,
		message.ResponseRulesCheckElapsed, message.HTTPElapsed,
		message.RequestSize, message.ResponseSize, message.RegisteredAt)

	return err
}
