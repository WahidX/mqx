package repository

import (
	"context"
	"database/sql"
	"mqx/internal/entities"
)

type Repository interface {
	EnqueueMessage(ctx context.Context, msg *entities.MessageRow) (int64, error)
	DequeueMessage(ctx context.Context, topic string) (*entities.MessageRow, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
