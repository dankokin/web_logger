package services

import (
	"strings"
	"time"

	"github.com/web_logger/models"
)

type DataStore interface {
	// Function that gets all logs from database by filters
	GetAllLogs(map[string]interface{}) ([]models.WAFMessage, error)
}

func (db *DB)GetAllLogs(filters map[string]interface{}) ([]models.WAFMessage, error) {
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("SELECT \"registered_at\", \"message\", \"target\" FROM public.\"Messages\" where \"registred_at\" = $1")

	args := make([]interface{}, 0, 2)

	var beginInterval time.Time
	if interval, exists := filters["interval"]; exists {
		beginInterval = time.Now().Add(interval.(time.Duration) * -1)
	} else {
		beginInterval = time.Now().Add(time.Hour * 24 * -1)
	}
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
