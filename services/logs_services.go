package services

import (
	"strings"
	"time"

	"github.com/dankokin/web_logger/models"
)

type DataStore interface {
	// Function that gets all logs from database by filters
	GetAllLogs(map[string]interface{}) ([]models.WAFMessage, error)

	// Simple function that saves WAFMessage to database
	SaveMessage(models.WAFMessage) error
}

func (db *DB)GetAllLogs(filters map[string]interface{}) ([]models.WAFMessage, error) {
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("SELECT \"registered_at\", \"message\", \"target\" FROM public.\"Messages\" where \"registred_at\" >= $1")

	args := make([]interface{}, 0, 2)

	interval := filters["interval"]
	beginInterval := time.Now().Add(-24 * time.Hour * time.Duration(interval.(int)))

	args = append(args, beginInterval)

	if target, exists := filters["target"]; exists {
		args = append(args, target)
		queryBuilder.WriteString(" and \"target\" = $2")
	}

	rows, err := db.Query(queryBuilder.String(), args...)
	if err != nil {
		return nil, err
	}

	messages := make([]models.WAFMessage, 0, 256)
	for rows.Next() {
		var msg models.WAFMessage
		err = rows.Scan(
			&msg.RegisteredAt,
			&msg.Message,
			&msg.Target)

		if err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func (db *DB)SaveMessage(message models.WAFMessage) error {
	_, err := db.Exec("INSERT INTO public.\"Messages\" \"registered_at\", \"message\", \"target\" VALUES ($1, $2, $3)",
		message.Message,
		message.RegisteredAt,
		message.Target)

	return err
}
