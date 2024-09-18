package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-mq/internal/entities"

	"go.uber.org/zap"
)

func (r *repository) DequeueMessage(ctx context.Context, topic string) (*entities.MessageRow, error) {
	var msg = &entities.MessageRow{}

	query := fmt.Sprintf(`
		DELETE FROM messages 
		WHERE id = (
			SELECT id 
			FROM messages 
			WHERE topic = '%s' 
			ORDER BY timestamp ASC 
			LIMIT 1
		) RETURNING data, timestamp, topic;
	`, topic)

	// Execute the query
	err := r.db.QueryRowContext(ctx, query).Scan(&msg.Data, &msg.Timestamp, &msg.Topic)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		zap.L().Warn("failed to execute query", zap.Error(err))
		return nil, fmt.Errorf("query execution failed: %v", err)
	}

	return msg, nil
}
