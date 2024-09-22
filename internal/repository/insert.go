package repository

import (
	"context"
	"mqx/internal/entities"

	"go.uber.org/zap"
)

func (r *repository) EnqueueMessage(ctx context.Context, msg *entities.MessageRow) (int64, error) {
	cmd, err := r.db.ExecContext(ctx, `INSERT INTO messages (data, timestamp, topic) VALUES ($1, $2, $3)`, string(msg.Data), msg.Timestamp, msg.Topic)
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
