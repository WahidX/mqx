package repository

import (
	"context"
	"fmt"
	"go-mq/internal/entities"
)

func (r *repository) DequeueMessage(ctx context.Context, topic string) (*entities.MessageRow, error) {
	var msg = &entities.MessageRow{}

	query := `DELETE FROM table_name WHERE condition = ? RETURNING data, timestamp, topic, partition`
	condition := fmt.Sprintf("id = (select id from messages where topic = %s order by timestamp asc limit 1)", topic)

	// Execute the query
	err := r.db.QueryRowContext(ctx, query, condition).Scan(&msg.Data, &msg.Timestamp, &msg.Topic, &msg.Partition)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}

	return msg, nil
}
