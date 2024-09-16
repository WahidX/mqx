package repository

import (
	"context"
	"go-mq/internal/entities"

	"go.uber.org/zap"
)

func (r *repository) PushMessage(ctx context.Context, msg *entities.MessageRow) (int64, error) {
	cmd, err := r.db.ExecContext(ctx, `INSERT INTO messages (data, timestamp, topic, partition) VALUES ($1, $2, $3, $4)`, string(msg.Data), msg.Timestamp, msg.Topic, msg.Partition)
	if err != nil {
		zap.L().Warn("failed to insert message", zap.Error(err))
		return 0, err
	}

	id, err := cmd.LastInsertId()
	if err != nil {
		zap.L().Warn("failed to get last insert id", zap.Error(err))
		return 0, err
	}

	return id, nil
}
