package repository

import (
	"context"
	"database/sql"
	"go-mq/internal/entities"
)

type Repository interface {
	PushMessage(ctx context.Context, msg *entities.MessageRow) (int64, error)
	// GetMessage(context.Context, string) (*entities.Message, error)

}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
